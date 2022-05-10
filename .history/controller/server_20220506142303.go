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

func NewServer() *server {
	s := &server{}
	s.Routes()
	return s
}

func (s *server) config(db *sql.DB, router *mux.Router) {
	s.db = db
	s.router = router
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
