package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// struct dengan json struct tag
// json struct tag memungkinkan kita membuat nama alias tipe data dengan nama yang berbeda dengan response API.
// lebih simpelnya mematakan response yang didapat dari source API ke bentuk yang kita inginkan.
type ResponseGetProductByID struct {
	Judul string `json:"title"`
	Kategori string `json:"category"`
	Deskripsi string `json:"description"`
}

func main(){
	// 1. Buat request
	req,err := http.NewRequest("GET","https://fakestoreapi.com/products/2",nil)
	// Fungsi http.NewRequest untuk membuat sebuah request ke API, fungsi ini memiliki 3 parameter
	// yaitu: method, url API, dan body
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// 2. Buat client
	client := http.Client{}
	// Client di atas bisa dipakai untuk mengirim request.

	// 3. Panggil request dengan client
	res,err := client.Do(req)
	// Client.Do adalah fugnsi untuk mengeksekusi request, fungsi ini mengebalikan response API dan error
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// tutup response body agar tidak membebani memory
	defer res.Body.Close()
	// Menjamin body selalu ditutup pada akhir fungsi, agar tidak bocor resource.

	// 4. Baca response body
	resBody, err := io.ReadAll(res.Body)
	// ReadAll berfugnsi untuk membaca response yang didapat dari API, fungsi ini mengembalikan 
	// data berupa berbentuk []byte dan error
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// 5. Convert ke tipe data custom ResponseGetProductByID
	var productResponse ResponseGetProductByID
	err = json.Unmarshal(resBody, &productResponse)
	// json.Unmarshal digunakan untuk memetakan data byte yang didapat di response ke tipe data custom atau struct
	// yang sudah dibuat atau ditentukan.
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("Title: ",productResponse.Judul)
	fmt.Println("Description: ",productResponse.Deskripsi)
	fmt.Println("Category: ",productResponse.Kategori)
}