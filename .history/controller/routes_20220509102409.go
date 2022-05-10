package controller

import (
	"log"
	"net/http"

	Router "github.com/gorilla/mux"
)

//set up all of the routes that the server will use
func (s *server) Routes() {
	s.router = Router.NewRouter()
	s.router.HandleFunc("/login", s.handleLogin()).Methods("POSTs")
	s.router.HandleFunc("/", s.handleHomePage()).Methods("GET")
	s.router.HandleFunc("/products/{product}", s.handleProducts()).Methods("GET")
	http.Handle("/", s.router)
	log.Fatalln(http.ListenAndServe(":10000", nil))
}
