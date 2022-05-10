package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type Server struct {
	db     *sql.DB
	router *mux.Router
}
