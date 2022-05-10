package controller

import (
	Router "github.com/gorilla/mux"
)

func routes() error {
	router := Router.NewRouter()
	router.HandleFunc("/login", handleLogin).Methods("GET")
	router.HandleFunc("/", handleHomePage).Methods("GET")
}
