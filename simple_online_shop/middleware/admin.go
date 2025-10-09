package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mengambil password admin dari ENV
		// disini password admin = "password"
		key := os.Getenv("ADMIN_SECRET_KEY")

		// 1. Ambil header authori zation
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.JSON(401, gin.H{"error": "maaf akses ini tidak diizinkan"})
			c.Abort() // Fungsi ini digunakan untuk menghentikan proses request
			return
		}

		// 2. Validasi header sesuai dengan password admin
		if auth != key {
			c.JSON(401, gin.H{"error": "maaf akses ini tidak diizinkan"})
			c.Abort() // Fungsi ini digunakan untuk menghentikan proses request
			return
		}

		// 3. Jika tidak ada error, maka proses request akan diteruskan
		c.Next()
	}
}
