package database

import (
	"fmt"
	"gin_backend/config"
	"gin_backend/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Get env value from .env
	dbUser := config.GetENV("DB_USER", "postgres")
	dbPass := config.GetENV("DB_PASS", "mysecretpassword")
	dbHost := config.GetENV("DB_HOST", "localhost")
	dbPort := config.GetENV("DB_PORT", "5432")
	dbName := config.GetENV("DB_Name", "sntr_coding_golang")

	// format URI untuk postgres
	URI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	// Koneksi ke Database
	var err error
	DB, err = gorm.Open(postgres.Open(URI), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connected successfully")

	// Migrate database model ke DB
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Database migrated successfully!")
}
