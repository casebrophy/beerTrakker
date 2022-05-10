package server

import (
	"fmt"
	"net/http"
)

func routes() error {
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/", handleHomePage)
}