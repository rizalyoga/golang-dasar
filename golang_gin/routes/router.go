package routes

import (
	"gin_backend/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Inisialisasi gin
	router := gin.Default()

	// Router register
	router.POST("/api/v1/register", controller.Register)

	// route login
	router.POST("/api/v1/login", controller.Login)

	return router
}
