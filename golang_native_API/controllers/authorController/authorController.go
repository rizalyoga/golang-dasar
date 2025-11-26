package authorcontroller

import (
	"encoding/json"
	"errors"
	"golang_native_api/database"
	"golang_native_api/models"
	"net/http"

	"gorm.io/gorm"
)



func Index(w http.ResponseWriter, r *http.Request) {
	var authors []models.Author
	w.Header().Set("Content-Type", "application/json")

	if err := database.DB.Find(&authors).Error; err != nil {
		// Periksa jika errornya adalah 'record not found'
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message":    "Data tidak ditemukan",
				"statusCode": http.StatusNotFound,
				"success":    false,
			})
			return
		}

		// Untuk semua error database lainnya
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":    "Gagal mengambil data dari database: " + err.Error(),
			"statusCode": http.StatusInternalServerError,
			"success":    false,
		})
		return
	}

	// Kirim response sukses
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":       authors,
		"message":    "Data berhasil diambil",
		"statusCode": http.StatusOK,
		"success":    true,
	})
}
