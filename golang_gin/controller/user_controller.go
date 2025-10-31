package controller

import (
	"gin_backend/database"
	"gin_backend/helpers"
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

	// Buat custom response sehingga password user tidak ikut menjadi respon API
	var userResponses []structs.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, structs.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	// Kirim data user yang sudah didapat dari DB
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Lists Data Users",
		Data:    userResponses,
	})
}

func CreateUser(c *gin.Context) {
	// struct user request
	var req = structs.UserCreateRequest{}

	// Bind JSON request ke struct UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Inisialisasi user baru
	user := models.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: helpers.HashPassword(req.Password),
	}

	// Simpan user ke database
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to create user",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Kirim respon sukses
	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "User created successfully",
		Data: structs.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}
