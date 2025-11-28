package bookcontroller

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
	var books []models.Book
	var BooksResponse []models.BooksResponse

	if err := database.DB.Joins("Author").Find(&books).Find(&BooksResponse).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, http.StatusNotFound, "Sorry, data book not found", nil)
			return
		}

		helper.Response(w, http.StatusInternalServerError, "Failed to retrieve data from database:"+err.Error(), nil)
		return
	}

	helper.Response(w, http.StatusOK, "Data successfully retrieved", BooksResponse)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	var author models.Author

	if err := database.DB.First(&author, book.AuthorID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, http.StatusNotFound, "Data author not found: "+err.Error(), nil)
			return
		}

		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if err := database.DB.Create(&book).Error; err != nil {
		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.Response(w, http.StatusCreated, "Book data created successfuly", nil)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	var bookResponse models.BooksResponse

	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	if err := database.DB.Joins("Author").First(&book, id).First(&bookResponse, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, http.StatusNotFound, "Data book not found: "+err.Error(), nil)
			return
		}

		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.Response(w, http.StatusOK, "Data book successfuly retrieved", bookResponse)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	if err := database.DB.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, http.StatusNotFound, "Data book not found: "+err.Error(), nil)
			return
		}

		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	var bookPayload models.Book
	if err := json.NewDecoder(r.Body).Decode(&bookPayload); err != nil {
		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	var author models.Author
	if err := database.DB.First(&author, bookPayload.AuthorID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, http.StatusNotFound, "Data author not found: "+err.Error(), nil)
			return
		}

		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if err := database.DB.Where("id = ?", id).Updates(&bookPayload).Error; err != nil {
		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.Response(w, http.StatusOK, "Book data successfuly updated", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	if err := database.DB.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, http.StatusNotFound, "Data book not found: "+err.Error(), nil)
			return
		}

		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if err := database.DB.Delete(&book, id).Error; err != nil {
		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.Response(w, http.StatusOK, "Data book successfuly deleted", nil)
}
