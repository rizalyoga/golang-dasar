package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// connURI := "postgres://postgres:mysecretpassword@localhost:5432/simple_online_shop?sslmode=disable"
	db, err := sql.Open("pgx", os.Getenv("DB_URI"))
	// fungsi os.Getenv = untuk mendapatkan nilai env.
	// untuk membuat env digolang bisa jalanlkan perintah berikut di terminal :
	// export Nama_ENV =  nilai_ENV
	if err != nil {
		fmt.Printf("Gagal membuat koneksi ke database! %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("Gagal ping ke database! %v\n", err)
		os.Exit(1)
	}
}
