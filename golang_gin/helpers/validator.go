package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// TranslateErrorMessage menangani validasi dari validator.v10 dan duplikasi entri dari GORM
func TranslateErrorMessage(err error) map[string]string {
	// Membuat map untuk menampung pesan error
	errorsMap := make(map[string]string)

	// Handle validasi dari validator.v10
	validationsError, ok := err.(validator.ValidationErrors)
	if ok {
		for _, fieldError := range validationsError {
			field := fieldError.Field() // Menyimpan nama field yang gagal validasi
			switch fieldError.Tag() {   // Menangani berbagai jenis validasi
			case "required":
				// / Pesan error jika field kosong
				errorsMap[field] = fmt.Sprintf("%s is required", field)
			case "email":
				// Pesan error jika format email tidak valid
				errorsMap[field] = "Invalid email format"
			case "unique":
				// Pesan error jika data sudah ada
				errorsMap[field] = fmt.Sprintf("%s duplicate key value violates unique constraint", field)
			case "min":
				// Pesan error jika nilai terlalu pendek
				errorsMap[field] = fmt.Sprintf("%s must be at least %s characters", field, fieldError.Param())
			case "max":
				// Pesan error jika nilai terlalu pendek
				errorsMap[field] = fmt.Sprintf("%s must be at most %s characters", field, fieldError.Param())
			case "numeric":
				// Pesan error jika nilai bukan angka
				errorsMap[field] = fmt.Sprintf("%s must be a number", field)
			default:
				// Pesan error default untuk kesalahan validasi lainnya
				errorsMap[field] = "Invalid value"
			}
		}
	}

	if err != nil {
		// Cek jika error mengandung "Duplicate entry" (duplikasi data di database)
		if strings.Contains(err.Error(), "duplicate key") {
			if strings.Contains(err.Error(), "username") {
				// Pesan error jika username sudah ada
				errorsMap["Username"] = "Username already exists"
			}
			if strings.Contains(err.Error(), "email") {
				// Pesan error jika email sudah ada
				errorsMap["Email"] = "Email already exists"
			}
		} else if err == gorm.ErrRecordNotFound {
			// Jika data yang dicari tidak ditemukan di database
			errorsMap["Error"] = "Record not found"
		}
	}

	// Return map yang berisi pesan error
	return errorsMap
}

// IsDuplicateEntryError mendeteksi apakah error dari database adalah duplicate entry
func IsDuplicateEntryError(err error) bool {
	// Cek apakah error merupakan duplikasi entri
	return err != nil && strings.Contains(err.Error(), "duplicate key")
}
