package helpers

import (
	"gin_backend/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Get JWT_SECRET value from .env
var jwtKey = []byte(config.GetENV("JWT_SECRET", "sNtr_c0D1n6_g0l4N6"))

func GenerateToken(username string) string {
	// Set waktu kadaluarsa token, disini token kadaluarsa adalah 60 menit
	expirationTime := time.Now().Add(60 * time.Minute)

	// Membuat calim jwt
	// Subject berisi username dan expiredAt yang menentukan expired token
	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	// Membuat token baru dengan klaim yang sudah dibuat
	// Membuat algoritma HS256 untuk tanda tangan token
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)

	// return token dalam bentuk string
	return token
}
