package authorcontroller

import (
	"encoding/json"
	"errors"
	"golang_native_api/database"
	"golang_native_api/helper"
	"golang_native_api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var authors []models.Author

	if err := database.DB.Find(&authors).Error; err != nil {
		// Periksa jika errornya adalah 'record not found' 404
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, http.StatusNotFound, "Sorry data not found", nil)
			return
		}

		// Untuk semua error database lainnya
		helper.Response(w, http.StatusInternalServerError, "Failed to retrieve data from database: "+err.Error(), nil)
		return
	}

	// Kirim response sukses
	helper.Response(w, http.StatusOK, "Data successfully retrieved", authors)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := database.DB.Create(&author).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Author data created successfuly", nil)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	if err := database.DB.First(&author, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, http.StatusNotFound, "Author data not found", nil)
			return
		}

		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.Response(w, http.StatusOK, "Detail author data successfully retrieved", author)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	if err := database.DB.First(&author, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, http.StatusNotFound, "Author data not found", nil)
			return
		}

		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := database.DB.Where("id = ?", id).Updates(&author).Error; err != nil {
		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.Response(w, http.StatusOK, "Author data successfuly update", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	if err := database.DB.First(&author, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, http.StatusNotFound, "Author data not found", nil)
			return
		}

		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if err := database.DB.Delete(&author, id).Error; err != nil {
		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.Response(w, http.StatusOK, "Author data deleted successfully", nil)
}
