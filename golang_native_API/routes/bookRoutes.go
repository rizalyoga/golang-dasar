package routes

import (
	bookcontroller "golang_native_api/controllers/bookController"
	"golang_native_api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func BookRoutes(r *mux.Router) {
	router := r.PathPrefix("/books").Subrouter()

	router.HandleFunc("", bookcontroller.Index).Methods("GET")
	router.Handle("", middlewares.BookValidatorMiddleware(http.HandlerFunc(bookcontroller.Create))).Methods("POST")
	router.HandleFunc("/{id}/detail", bookcontroller.Detail).Methods("GET")
	router.Handle("/{id}/update", middlewares.BookValidatorMiddleware(http.HandlerFunc(bookcontroller.Update))).Methods("PUT")
	router.HandleFunc("/{id}/delete", bookcontroller.Delete).Methods("DELETE")
}
