package main

import (
	"gin_backend/config"
	"gin_backend/database"
	"gin_backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Jalankan fungsi LoadENV dari package config
	config.LoadENV()

	// Load database
	database.InitDB()

	// setup router
	router := routes.SetupRouter()

	// Routing menggunakan method GET
	router.GET("/", func(ctx *gin.Context) {

		// Return JSON untuk route "/"
		ctx.JSON(200, gin.H{
			"message": "Hello, Backend already running in port :" + config.GetENV("APP_PORT", "5500"),
		})
	})

	// Run server di port 5500
	router.Run(":" + config.GetENV("APP_PORT", "5000"))
}
