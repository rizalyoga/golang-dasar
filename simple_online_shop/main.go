package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"simple_online_shop/handler"
	"simple_online_shop/middleware"

	"github.com/gin-gonic/gin"
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

	// create table in db
	_, err = migrate(db)
	if err != nil {
		fmt.Printf("Gagal melakukan migrasi database: %v\n", err)
		os.Exit(1)
	}

	// create server
	r := gin.Default()

	r.GET("api/v1/products", handler.ListProducts(db))
	r.GET("api/v1/products/:id", handler.GetProductsById(db))
	r.POST("api/v1/checkout")

	r.POST("api/v1/orders/:id/confirm")
	r.GET("api/v1/orders/:id")

	r.POST("api/v1/admin/products", middleware.AdminOnly(), handler.CreateProduct(db))
	r.PUT("api/v1/admin/products/:id", middleware.AdminOnly(), handler.UpdateProduct(db))
	r.DELETE("api/v1/admin/products/:id", middleware.AdminOnly(), handler.DeleteProduct(db))

	server := &http.Server{
		Addr:    ":5500",
		Handler: r,
	}

	err = server.ListenAndServe()
	if err != nil {
		fmt.Printf("Gagal menjalankan server: %v\n", err)
		os.Exit(1)
	}
}
