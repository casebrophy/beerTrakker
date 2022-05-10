package controller

import (
	"log"
	"net/http"

	Router "github.com/gorilla/mux"
)

func (server) Routes() {
	router := Router.NewRouter()
	router.HandleFunc("/login", handleLogin).Methods("GET")
	router.HandleFunc("/", handleHomePage).Methods("GET")
	router.HandleFunc("/products/{product}", handleProducts).Methods("GET")
	http.Handle("/", router)
	log.Fatalln(http.ListenAndServe(":10000", nil))
}
