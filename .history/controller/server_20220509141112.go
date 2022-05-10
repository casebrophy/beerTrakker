package controller

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type server struct {
	db     *sql.DB
	router *mux.Router
	store  *sessions.CookieStore
}

//initalizes server and all of its router
func NewServer() *server {
	s := &server{}
	s.Routes()
	return s
}

//sets the database
//configures the session store with private key
//TODO add other config parameters
func (s *server) Config(db *sql.DB) {
	s.db = db
	s.store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
}

//allows the server to work as a request handler
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
