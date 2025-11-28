package routes

import "github.com/gorilla/mux"

func IndexRoutes(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()

	AuthorRoutes(api)
	BookRoutes(api)
}
