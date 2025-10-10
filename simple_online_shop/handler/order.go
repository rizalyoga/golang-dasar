package handler

import (
	"database/sql"
	"log"
	"math/rand"
	"simple_online_shop/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Checkout Order
func CheckoutOrder(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil data pesanan dari body
		var checkoutOrder model.Checkout
		// BindJSON = dekode data JSON yang diterima dari permintaan HTTP atau sumber lain menjadi sebuah struct (tipe data komposit) dalam Go
		err := c.BindJSON(&checkoutOrder)
		if err != nil {
			log.Printf("Terjadi kesalahan saat membaca request body %v\n", err)
			c.JSON(400, gin.H{"error": "Data product tidak valid"})
			return
		}

		// list id
		ids := []string{}
		// order quantity
		orderQty := make(map[string]int32)
		for _, o := range checkoutOrder.Products {
			ids = append(ids, o.ID)     // memasukan daftar id dari product yg dicekout dari req body user
			orderQty[o.ID] = o.Quantity // memasukan quantity order dari product
		}

		// Ambil poduk data dari DB
		// products berisi list data product berdasarkan id yang berada di ids
		products, err := model.SelectProductIn(db, ids)
		if err != nil {
			log.Printf("Terjadi kesalahan saat mengambil product %v\n", err)
			c.JSON(500, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		// Buat kata sandi yang nantinya akan di tampilkan diresponse
		passcode := generatePasscode(5)

		// Hash kata sandi yang didapatkan dari generatePasscode
		hashCode, err := bcrypt.GenerateFromPassword([]byte(passcode), 10)
		if err != nil {
			log.Printf("Terjadi kesalahan saat hash password %v\n", err)
			c.JSON(500, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		// Conversi ke string dari []byte
		hashCodeString := string(hashCode)

		// Buar order & detail
		order := model.Order{
			ID:         uuid.New().String(),
			Email:      checkoutOrder.Email,   // checkoutOrder adalah data dari req body
			Address:    checkoutOrder.Address, // checkoutOrder adalah data dari req body
			Passcode:   &hashCodeString,
			GrandTotal: 0,
		}

		details := []model.OrderDetail{}
		// Looping data products yang didapat dari fungsi SelectProductIn
		for _, p := range products {
			total := p.Price * int64(orderQty[p.ID])

			detail := model.OrderDetail{
				ID:        uuid.New().String(),
				OrderID:   order.ID,
				ProductID: p.ID,
				Quantity:  orderQty[p.ID],
				Price:     p.Price,
				Total:     total,
			}

			details = append(details, detail)

			order.GrandTotal += total
		}

		// Membuat data Order dan Detail Order di DB
		err = model.CreateOrder(db, order, details)
		if err != nil {
			log.Printf("Terjadi kesalahan saat membuat order %v\n", err)
			c.JSON(500, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		orderWithDetail := model.OrderWithDetail{
			Order:   order,
			Details: details,
		}

		// Merubah passcode pada Order response menjadi nilai asli passcode bukan nilai hash
		orderWithDetail.Order.Passcode = &passcode

		c.JSON(200, orderWithDetail)
	}
}

// Fungsi untuk generate password secara acak berdasarkan karakter yang sudah didaftar
func generatePasscode(length int) string {
	// Daftar karakter yang nantinya akan diambil secara acak sebanyak nilai parameter length
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	code := make([]byte, length)
	for i := range code {
		code[i] = charset[randomGenerator.Intn(len(charset))]
	}

	return string(code)
}

// Fungsi untuk confirm order
func ConfirmOrder(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil id order dari param
		id := c.Param("id")

		// Baca request body
		var confirmReq model.Confirm
		err := c.BindJSON(&confirmReq)
		if err != nil {
			log.Printf("Terjadi kesalahan saat membaca req body: %v\n", err)
			c.JSON(400, gin.H{"error": "Data pesanan tidak valid"})
			return
		}

		// ambil data order dari DB
		order, err := model.SelectOrderById(db, id)
		if err != nil {
			log.Printf("Terjadi kesalahan saat membaca data pesanan: %v\n", err)
			c.JSON(500, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		if order.Passcode == nil {
			log.Printf("Terjadi kesalahan saat membaca data pesanan: %v\n", err)
			c.JSON(500, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		// cocokan passcode pesanan
		err = bcrypt.CompareHashAndPassword([]byte(*order.Passcode), []byte(confirmReq.Passcode))
		if err != nil {
			log.Printf("Terjadi kesalahan saat mencocokan password: %v\n", err)
			c.JSON(401, gin.H{"error": "Tidak diizinkan mengakses pesanan"})
			return
		}

		// pasktikan pesanan belum dibayar
		if order.PaidAt != nil {
			log.Println("Pesanan sudah dibayar")
			c.JSON(400, gin.H{"error": "Pesanan sudah dibayar"})
			return
		}

		// cocokan jumlah pembayaran
		if order.GrandTotal != confirmReq.Amount {
			log.Printf("Jumlah harga tidak sesuai: %v\n", confirmReq.Amount)
			c.JSON(400, gin.H{"error": "Jumlah pembayaran tidak sesuai"})
			return
		}

		// update informasi pesanan
		current := time.Now()
		err = model.UpdateOrderByID(db, id, confirmReq, current)
		if err != nil {
			log.Printf("Terjadi kesalahan saat memperbarui data pesanan: %v\n", err)
			c.JSON(400, gin.H{"error": "Data pesanan tidak valid"})
			return
		}

		order.Passcode = nil

		order.PaidAt = &current
		order.PaidBank = &confirmReq.Bank
		order.PaidAccount = &confirmReq.AccountNumber

		c.JSON(200, order)
	}
}

func GetOrder(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get params id
		id := c.Param("id")

		// ambil passcode dari query param url
		passcode := c.Query("passcode")

		// ambil data order dari DB
		order, err := model.SelectOrderById(db, id)
		if err != nil {
			log.Printf("Terjadi kesalahan saat mengambil data pesanan: %v\n", err)
			c.JSON(500, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		if order.Passcode == nil {
			log.Println("Passcode tidak valid")
			c.JSON(500, gin.H{"error": "Passcode tidak valid"})
			return
		}

		// cocokan kata santi / passcode
		err = bcrypt.CompareHashAndPassword([]byte(*order.Passcode), []byte(passcode))
		if err != nil {
			log.Printf("Terjadi kesalahan saat mencocokan password: %v\n", err)
			c.JSON(401, gin.H{"error": "Tidak diizinkan mengakses pesanan"})
			return
		}

		order.Passcode = nil

		c.JSON(200, order)
	}
}
