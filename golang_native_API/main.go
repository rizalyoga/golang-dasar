package main

import (
	"fmt"
	"golang_native_api/config"
	"golang_native_api/database"
	"golang_native_api/routes"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadENV()
	database.InitDB()

	r := mux.NewRouter()
	routes.IndexRoutes(r)

	log.Println("Server running in port", config.GetENV("APP_PORT", "5500"))
	http.ListenAndServe(fmt.Sprintf(":%v", config.GetENV("APP_PORT", "5500")), r)
}
