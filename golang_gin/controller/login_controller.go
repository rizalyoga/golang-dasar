package controller

import (
	"gin_backend/database"
	"gin_backend/helpers"
	"gin_backend/models"
	"gin_backend/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var req = structs.UserLoginRequest{}
	var user = models.User{}

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
	}

	// Cari user berdasarkan username yang diberikan di database
	err = database.DB.Where("username = ?", req.Username).First(&user).Error
	// Jika tidak ditemukan, kirimkan respons error Unauthorized
	if err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "User Not Found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Bandingkan password yang dimasukkan dengan password yang sudah di-hash di database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	// Jika tidak cocok, kirimkan respons error Unauthorized
	if err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "Invalid Password",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Jika login berhasil , generate token untuk user
	token := helpers.GenerateToken(user.Username)

	// Kirimkan response sukses dengan status OK dan data user serta token
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Login Success",
		Data: structs.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
			Token:     &token,
		},
	})
}
