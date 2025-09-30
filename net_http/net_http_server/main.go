package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Products struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
}

func main(){
	// 1. buat route multiplexer
	mux := http.NewServeMux()

	// 3. tambahkan handler ke mux
	// Mux adalah module pada golang yang berfugnsi untuk routing
	mux.HandleFunc("GET /", indexHandler)
	mux.HandleFunc("GET /products", listProduct)
	mux.HandleFunc("POST /products", createProduct)
	mux.HandleFunc("PUT /products/{id}", updateProduct)
	mux.HandleFunc("DELETE /products/{id}", deleteProduct)

	// 4. buat server
	server := http.Server{
		Handler: mux,
		// Handler disini untuk menentukan routing yang dipakai server
		Addr: ":5500",
		// Addr disini adalah address, yang mana digunakan untuk menentukan port mana yang diguakan oleh server.
	}

	// 5. jalankan server
	fmt.Println("Server run in port 5500")
	server.ListenAndServe()
}

var database = map[int]Products{}
var lastID int = 0

// 2. fungsi handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	var messages string = "Welcome in golang dasar CRUD Products"
	w.Write([]byte(messages))
}

func listProduct(w http.ResponseWriter,r *http.Request){
	// slice untuk response
	var products []Products

	// Iterasi untuk menambahkan nilai products pada map ke slice products
	for _,v := range database {
		products = append(products,v)
	}

	// Ubah data menjadi json
	data,err := json.Marshal(products)
	// json.Marshal digunakan untuk mengonversi data dari tipe Golang (seperti struct, map, atau slice) menjadi format
	// JSON dalam bentuk byte slice (irisan byte), yang kemudian bisa dikonversi menjadi string untuk dikirim atau
	// disimpan.

	// Mengirim response json ke client
	if err != nil {
		w.Header().Set("Content-Type","application/json") // Header digunakan untuk mengirim tipe repsponse
		w.WriteHeader(500) // WriteHeader untuk mengirim status code
		w.Write([]byte("Terjadi Kesalahan")) // Write untuk mengirim body response, karena tidak bisa mengirim string, maka menggunakan byte
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(200)
	w.Write(data)
}

func createProduct(w http.ResponseWriter,r *http.Request){
	// Mendapatkan nilai body dalam bentuk byte
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(400)
		w.Write([]byte("Kesalahan pada request"))
		return
	}

	var products Products
	err = json.Unmarshal(bodyByte, &products)
	// Mengonversi JSON dalam byte ke struct, dan nilai products akan bernilai sesuai data yang didapat setelah
	// konversi
	if err != nil {
		w.Header().Set("Content-type","application/json")
		w.WriteHeader(400)
		w.Write([]byte("Kesalahan dalam request"))
		return
	}

	// Tambah nilai lastID untuk id product
	lastID++
	products.ID = lastID

	// Menambahkan products ke map database
	database[products.ID] = products

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(201)
	w.Write([]byte("Products baru berhasil dibuat"))
}

func updateProduct(w http.ResponseWriter,r *http.Request){
	// Mendapatkan nilai params dari URL
	productID := r.PathValue("id")

	// Convert string ke int
	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(404)
		w.Write([]byte("Product ID tidak ditemukan"))
		return
	}

	// Baca / mendapatkan nilai body
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(400)
		w.Write([]byte("Kesalahan pada request"))
		return
	}

	var products Products
	// Mengonversi JSON dalam byte ke struct, dan nilai products akan bernilai sesuai data yang didapat setelah
	// konversi
	err = json.Unmarshal(bodyByte, &products)
	if err != nil {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(400)
		w.Write([]byte("Kesalahan pada request"))
		return
	}

	// membuat nilai products.id sesuai dengan nilai product id di params url 
	products.ID =  productIDInt

	// update nilai products di database sesuai dengan nilai products dari body request
	database[productIDInt] = products

	// Ubah data menjadi json
	data,err := json.Marshal(products)
	if err != nil {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(500)
		w.Write([]byte("Kesalahan pada pada server"))
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(200)
	w.Write(data)
}

func deleteProduct(w http.ResponseWriter,r *http.Request){
	// Mendapatkan nilai params dari url
	productID := r.PathValue("id")

	// Convert string ke int
	productIDInt ,err := strconv.Atoi(productID)
	if err != nil {
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(404)
		w.Write([]byte("Data tidak ditemukan"))
	}

	// Hapus data dari map
	delete(database, productIDInt)

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(204)
}