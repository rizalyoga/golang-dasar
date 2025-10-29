package middlewares

import (
	"fmt"
	"gin_backend/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Mengambil value JWT_SECRET dari .env, jika tidak ada akan menggunakan "secret_key"
var jwtKey = []byte(config.GetENV("JWT_SECRET", "sNtr_c0D1n6_g0l4N6"))

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Ambil header authorization dari request
		tokenString := c.GetHeader("Authorization")
		// Cek error
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token is required",
			})

			c.Abort() // Menghentikan request selanjutnya
			return
		}

		// Menhapus prefix atau kata "Bearer" dari token
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Membuat struct untuk menampung klaim token
		claims := &jwt.RegisteredClaims{}

		// Parse token dan verifikasi dengan jwtKey
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			// Kembalikan kunci rahasia untuk memverifikasi token
			return jwtKey, nil
		})

		fmt.Printf("Nilai token: %v\n:", token)
		fmt.Printf("Nilai claims: %v\n:", claims)

		//  Jika token tidak valid atau terjadi error saat parsing
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})

			c.Abort() // Hentikan request berikutnya
			return
		}

		// Simpan klaim "sub" ke dalam context
		c.Set("username", claims.Subject)
		// Lanjut ke header berikutnya
		c.Next()
	}
}
