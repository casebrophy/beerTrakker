package controller

import (
	"fmt"
	"net/http"
)

type User struct {
	id       int
	uName    string
	password string
}

func handleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "login")
	}
}
