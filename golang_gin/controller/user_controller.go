package controller

import (
	"gin_backend/database"
	"gin_backend/models"
	"gin_backend/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindUser(c *gin.Context) {
	// Create variable slace user untuk menampung data user
	var users []models.User

	// Ambil data user dari database
	database.DB.Find(&users)

	// Kirim data user yang sudah didapat dari DB
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Lists Data Users",
		Data:    users,
	})

}
