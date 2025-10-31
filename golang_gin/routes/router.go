package routes

import (
	"gin_backend/controller"
	"gin_backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Inisialisasi gin
	router := gin.Default()

	// Router register
	router.POST("/api/v1/register", controller.Register)

	// route login
	router.POST("/api/v1/login", controller.Login)

	// route find users
	router.GET("/api/v1/users", middlewares.AuthMiddleware(), controller.FindUser)

	// route user create
	router.POST("/api/v1/users", middlewares.AuthMiddleware(), controller.CreateUser)

	return router
}
