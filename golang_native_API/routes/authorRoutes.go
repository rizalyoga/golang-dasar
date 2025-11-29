package routes

import (
	authorcontroller "golang_native_api/controllers/authorController"
	validators "golang_native_api/middlewares" // Import the new middleware
	"net/http"                                 // Import http package for http.HandlerFunc

	"github.com/gorilla/mux"
)

func AuthorRoutes(r *mux.Router) {
	router := r.PathPrefix("/authors").Subrouter()

	router.HandleFunc("", authorcontroller.Index).Methods("GET")
	// Apply the middleware to Create and Update routes
	router.Handle("", validators.AuthorValidatorMiddleware(http.HandlerFunc(authorcontroller.Create))).Methods("POST")
	router.HandleFunc("/{id}/detail", authorcontroller.Detail).Methods("GET")
	router.Handle("/{id}/update", validators.AuthorValidatorMiddleware(http.HandlerFunc(authorcontroller.Update))).Methods("PUT")
	router.HandleFunc("/{id}/delete", authorcontroller.Delete).Methods("DELETE")
}
