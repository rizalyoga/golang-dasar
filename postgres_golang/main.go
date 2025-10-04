package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Untuk bisa menggunakan PostgreSQL di Golang, dibutuhkan 2 hal:
// - Modul database/sql (https://pkg.go.dev/database/sql)
// - Driver database (https://go.dev/wiki/SQLDrivers)

// Modul database/sql berisi interface yang mendefinisikan fungsi-fungsi yang bisa digunakan untuk bekerja dengan database relasional SQL.
// Untuk bisa menjalankan fungsi-fungsi tersebut di database tertentu, modul ini membutuhkan driver.

// Driver database mengimplementasikan interface yang didefinisikan dalam modul database/sql dan menjalankan fungsi tersebut ke sebuah system database seperti PostgreSQL.
// Driver PostgreSQL yang populer digunakan adalah pq dan pgx, namun driver pq sudah tidak lagi dikembangkan.


type Produk struct {
	ID uint
	Nama string
	Kategori string
	Harga int
}

// Fungsi - fungsi yang ada di modul database golang
// sql.Open Membuat koneksi ke database, memiliki return berupa sql.DB.
// DB.Close Menutup koneksi yang dibuat
// DB.Ping Memverifikasi status koneksi dan menghubungkan jika belum terhubung di sql.Open()
// DB.Query Menjalankan query yang memiliki return beberapa baris data, misal SELECT.
// DB.QueryRow Menjalankan query yang memiliki return 1 baris data, misal SELECT, INSERT dengan RETURNING.
// DB. Exec Menjalankan query tanpa return, misal INSERT, UPDATE, dan DELETE.

func main()  {
	// Download driver pgx dengan perintah "go get github.com/jackc/pgx/v5/stdlib"
	
	// postgresql://username:password@hostname:port/db_name?sslmode=disable/enable	
	// sslmode tergantung apakah menggnakan https atau http, jika menggunakan http/local maka nilai = disable.
	connURI := "postgresql://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable"

	// sql.Open memiliki 2 parameter = nama driver dan nama data source
	db, err := sql.Open("pgx",connURI)
	if err != nil {
		fmt.Printf("Gagal menghubungkan ke database: \n %v \n", err)
		os.Exit(1)
	}

	// db.Close fungsi untuk menutup koneksi
	defer db.Close()

	// db.ping berguna untuk verifikasi koneksi dengan database
	err = db.Ping()
	if err != nil {
		fmt.Printf("Terjadi kesalahan: \n %v \n", err)
		os.Exit(1)
	}

	fmt.Println("Database berhasil terhubung")

	// Membuat table produk
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS produk (
		id SERIAL PRIMARY KEY,
		nama VARCHAR(255),
		kategori VARCHAR(50),
		harga INT
	)`)

	if err != nil {
		fmt.Printf("Gagal membuat table database %v \n", err)
		os.Exit(1)
	}

	fmt.Println("Database berhasil dibuat...")
	
	// Proses insert data
	// _, err = db.Exec(`INSERT INTO produk (nama, kategori, harga) VALUES ($1,$2,$3)`,
	// "Kertas A4", "Kertas", 2000,
	// "Kertas A5", "Kertas", 1500,
	// "Buku 58", "Buku Tulis", 4500,
	// )
	// Tanda dolar ($) pada query di postgres disebut place holder
	// if err != nil {
	// 	fmt.Printf("Gagal membuat table database %v \n", err)
	// 	os.Exit(1)
	// }
	
	// fmt.Println("Data produk berhasil dibuat...") 

	// Proses get one data
	row := db.QueryRow(`SELECT id, nama, kategori, harga FROM produk WHERE id=$1`, 1)
	if row == nil {
		fmt.Printf("Gagal mengambil table database %v \n", err)
		os.Exit(1)
	}

	var produk Produk
	// row.Scan() = Menyimpan nilai dari sebuah kolom ke sebuah variable di Golang
	err = row.Scan(&produk.ID, &produk.Nama, &produk.Kategori, &produk.Harga)
	if err != nil {
		fmt.Printf("Gagal mengambil data: %v \n", err)
		os.Exit(1)
	}
	fmt.Println(produk)

	// Proses get all data
	rows, err := db.Query(`SELECT id, nama, kategori, harga FROM produk`)
	if err != nil || rows == nil{
		fmt.Printf("Gagal mengambil data: %v \n", err)
		os.Exit(1)
	}

	var produkSlice []Produk
	for rows.Next() {
		var produk Produk
		err =  rows.Scan(&produk.ID, &produk.Nama, &produk.Kategori, &produk.Harga)
		if err != nil {
			fmt.Printf("Gagal scan data: %v \n", err)
			os.Exit(1)
		}

		produkSlice = append(produkSlice, produk)
	}

	fmt.Println(produkSlice)	

	// Proses update data
	// _, err = db.Exec(`UPDATE produk SET nama=$1, kategori=$2, harga=$3 WHERE id=$4`, "Bakwan","Makanan",2000, 8)
	// if err != nil {
	// 	fmt.Printf("Gagal update data: %v \n", err)
	// 	os.Exit(1)
	// }

	// fmt.Println("Berhasil update data...")

	// Proses menghapus data
	_, err = db.Exec(`DELETE FROM produk WHERE id=$1`, 8)
	if err != nil {
		fmt.Printf("Gagal menghapus data: %v \n", err)
		os.Exit(1)
	}

	fmt.Println("Data berhasil di hapus")

	// Transaction
	// Fungsi - fungsi yang ada di transaction:
	//  DB.Begin Memulai transaction, memiliki return berupa sql.Tx.
	//  Tx.Commit Menyimpan hasil transaction
	//  Tx.Rollback Membatalkan transaction
	//  Tx.Query Sama dengan DB.Query() namun dilakukan dalam transaction
	//  Tx.QueryRow Sama dengan DB.QueryRow() namun dilakukan dalam transaction
	//  Tx.Exec Sama dengan DB.Exec() namun dilakukan dalam transaction


	// db.Begin(),Memulai transaction, memiliki return berupa sql.Tx.
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("Gagal membuat transaction, Error: %v \n", err)
		os.Exit(1)
	}

	// tx.Exec sama dengan db.Exec digunakan untuk melakukan query yang tidak memiliki return data
	_,err = tx.Exec(`DELETE FROM produk WHERE id=$1`, 12)
	if err != nil {
		fmt.Printf("Gagal menghapus data, Error: %v \n", err)
		// tx.Roleback() membatalkan transaction
		tx.Rollback()
		os.Exit(1)
	}

	// tx.Commit Menyimpan transaction
	tx.Commit()
	fmt.Println("Data berhasil di hapus...")
}
