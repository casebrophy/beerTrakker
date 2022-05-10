package controller

import (
	"log"
	"net/http"

	Router "github.com/gorilla/mux"
)

//set up all of the routes that the server will use
func (s *Server) Routes() {
	s.router = Router.NewRouter()
	s.router.HandleFunc("/login", s.handleLogin()).Methods("POST")
	s.router.HandleFunc("/login", s.handleLoginPage()).Methods("GET")
	s.router.HandleFunc("/", s.handleHomePage()).Methods("GET")
	s.router.HandleFunc("/products/{product}", s.handleProducts()).Methods("GET")
	http.Handle("/", s.router)
	log.Fatalln(http.ListenAndServe(":10000", nil))
}
