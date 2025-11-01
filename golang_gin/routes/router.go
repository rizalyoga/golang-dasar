package routes

import (
	"gin_backend/controller"
	"gin_backend/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Inisialisasi gin
	router := gin.Default()

	// Setup cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	// Router register
	router.POST("/api/v1/register", controller.Register)

	// route login
	router.POST("/api/v1/login", controller.Login)

	// route find users
	router.GET("/api/v1/users", middlewares.AuthMiddleware(), controller.FindUser)

	// route user create
	router.POST("/api/v1/users", middlewares.AuthMiddleware(), controller.CreateUser)

	// route get user by id
	router.GET("/api/v1/users/:id", middlewares.AuthMiddleware(), controller.FindUserById)

	// route update user data
	router.PUT("/api/v1/users/:id", middlewares.AuthMiddleware(), controller.UpdateUser)

	// route update user data
	router.DELETE("/api/v1/users/:id", middlewares.AuthMiddleware(), controller.DeleteUser)

	return router
}
