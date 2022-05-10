package controller

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	db     *sql.DB
	router *mux.Router
}

//initalizes server and all of its router
func NewServer() *server {
	s := &server{}
	s.Routes()
	return s
}

//sets the database
//TODO add other config parameters
func (s *server) Config(db *sql.DB) {
	s.db = db
}

//allows the server to work as a request handler
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
