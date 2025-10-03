package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Produk struct {
	ID uint
	Nama string
	Kategori string
	Harga int
}

func main()  {
	// postgresql://username:password@hostname:port/db_name?sslmode=disable/enable
	// sslmode tergantung apakah menggnakan https atau http, jika menggunakan http/local maka nilai = disable.
	connURI := "postgresql://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable"

	// sql.Open memiliki 2 parameter = napa driver dan nama data source
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
	_, err = db.Exec(`INSERT INTO produk (nama, kategori, harga) VALUES ($1,$2,$3)`,
	// "Kertas A4", "Kertas", 2000,
	// "Kertas A5", "Kertas", 1500,
	"Buku 58", "Buku Tulis", 4500,
	)
	// Tanda dolar ($) pada query di postgres disebut place holder
	if err != nil {
		fmt.Printf("Gagal membuat table database %v \n", err)
		os.Exit(1)
	}
	fmt.Println("Data produk berhasil dibuat...")

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
	_, err = db.Exec(`UPDATE produk SET nama=$1, kategori=$2, harga=$3 WHERE id=$4`, "Bakwan","Makanan",2000, 8)
	if err != nil {
		fmt.Printf("Gagal update data: %v \n", err)
		os.Exit(1)
	}

	fmt.Println("Berhasil update data...")

	// Proses menghapus data
	_, err = db.Exec(`DELETE FROM produk WHERE id=$1`, 8)
	if err != nil {
		fmt.Printf("Gagal menghapus data: %v \n", err)
		os.Exit(1)
	}

	fmt.Println("Data berhasil di hapus")

}
