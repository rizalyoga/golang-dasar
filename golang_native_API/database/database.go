package database

import (
	"fmt"
	"golang_native_api/config"
	"golang_native_api/models"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbUser := config.GetENV("DB_USER", "postgres")
	dbPass := config.GetENV("DB_PASSWORD", "admin")
	dbHost := config.GetENV("DB_HOST", "localhost")
	dbName := config.GetENV("DB_NAME", "db_book")
	dbPort := config.GetENV("DB_PORT", "1234")

	URI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	var err error

	db, err := gorm.Open(postgres.Open(URI), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&models.Author{}, &models.Book{})
	if err != nil {
		log.Fatal("Failed to migrate Author table", err)
	}

	DB = db
	log.Println("Database connected")
}
