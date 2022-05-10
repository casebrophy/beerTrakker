package controller

import (
	"log"
	"net/http"

	Router "github.com/gorilla/mux"
)

func Routes() {
	router := Router.NewRouter()
	router.HandleFunc("/login", handleLogin).Methods("GET")
	router.HandleFunc("/", handleHomePage).Methods("GET")
	log.Fatalln(http.ListenAndServe(":10000", nil))
}
