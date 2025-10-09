package handler

import (
	"database/sql"
	"errors"
	"log"
	"simple_online_shop/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ListProducts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil dari database
		products, err := model.SelectProducts(db)

		// berikan response
		if err != nil {
			log.Printf("Terjadi kesalahan saat mengambil data pada db: %v\n", err)
			c.JSON(500, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		responseData := map[string]any{
			"status_code": 200,
			"status":      "success",
			"message":     "Data berhasil didapatkan",
			"data":        products,
		}
		c.JSON(200, responseData)
	}
}

func GetProductsById(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// baca id dari params
		id := c.Param("id")

		// ambil data dari database
		product, err := model.SelectProductById(db, id)

		// beri response
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				log.Printf("Terjadi kesalahan saat mengambil data produk pada db: %v\n", err)
				c.JSON(500, gin.H{"error": "Data tidak ditemukan"})
				return
			}

			log.Printf("Terjadi kesalahan saat mengambil data pada db: %v\n", err)
			c.JSON(500, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		responseData := map[string]any{
			"status_code": 200,
			"status":      "success",
			"message":     "Data produk berhasil di dapatkan",
			"data":        product,
		}

		c.JSON(200, responseData)
	}
}

// Create Product
func CreateProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product model.Product

		// Membaca  request body
		err := c.Bind(&product) // Bind digunakan untuk membaca request body
		if err != nil {
			log.Printf("Terjadi kesalahan saat membaca req body %v\n", err)
			c.JSON(400, gin.H{"error": "Data produk tidak valid"})
			return
		}

		// Membuat ID untuk product
		product.ID = uuid.New().String()

		// Membuat produk data di DB
		err = model.InsertProduct(db, product)
		if err != nil {
			log.Printf("Terjadi kesalahan saat membuat produk %v\n", err)
			c.JSON(500, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		// Membuat reponse
		responseData := map[string]any{
			"status_code": 201,
			"status":      "success",
			"message":     "Data produk baru berhasil dibuat",
			"data":        product,
		}

		// Return response ke user
		c.JSON(201, responseData)
	}
}

// Update product
func UpdateProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Membaca id dari parameter
		id := c.Param("id")

		// Mendapatkan req body
		var product model.Product
		err := c.Bind(&product)
		if err != nil {
			log.Printf("Terjadi kesalahan saat membaca req body %v\n", err)
			c.JSON(400, gin.H{"error": "Data produk tidak valid"})
			return
		}

		// Mengambil data yang sudah ada dari DB
		productExisting, err := model.SelectProductById(db, id)
		if err != nil {
			log.Printf("Terjadi kesalahan saat mengambil produk %v\n", err)
			c.JSON(500, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		// Cek nilai produk yang baru
		if product.Name != "" {
			productExisting.Name = product.Name
		}

		if product.Price != 0 {
			productExisting.Price = product.Price
		}

		// Update data produk baru ke DB
		err = model.UpdateProduct(db, productExisting)
		if err != nil {
			log.Printf("Terjadi kesalahan saat update data produk %v\n", err)
			c.JSON(500, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		// Membuat response
		responseData := map[string]any{
			"status_code": 200,
			"status":      "success",
			"message":     "Data produk baru berhasil diperbarui",
			"data":        product,
		}

		// Return response ke user
		c.JSON(200, responseData)
	}
}

// Delete product
func DeleteProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get id param
		id := c.Param("id")

		// Delete produk di DB
		err := model.DeleteProduct(db, id)
		if err != nil {
			log.Printf("Terjadi kesalahan saat delete data produk %v\n", err)
			c.JSON(500, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		// Membuat response
		responseData := map[string]any{
			"status_code": 200,
			"status":      "success",
			"message":     "Data produk berhasil dihapus",
		}

		// Return response ke user
		c.JSON(200, responseData)
	}
}
