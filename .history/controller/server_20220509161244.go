package controller

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

type Server struct {
	db     *sql.DB
	router *mux.Router
	store  *sessions.CookieStore
}

//initalizes server and all of its router
func NewServer() *Server {
	s := &Server{}
	s.Routes()
	return s
}

//sets the database
//configures the session store with private key
//TODO add other config parameters
func (s *Server) Config(db *sql.DB) {
	s.db = db
	s.store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32))
}

//allows the server to work as a request handler
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
