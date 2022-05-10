package controller

import (
	"log"
	"net/http"

	Router "github.com/gorilla/mux"
)

func (s *server) Routes() {
	s.router = Router.NewRouter()
	s.router.HandleFunc("/login", handleLogin()).Methods("GET")
	s.router.HandleFunc("/", s.handleHomePage()).Methods("GET")
	s.router.HandleFunc("/products/{product}", handleProducts()).Methods("GET")
	http.Handle("/", s.router)
	log.Fatalln(http.ListenAndServe(":10000", nil))
}
