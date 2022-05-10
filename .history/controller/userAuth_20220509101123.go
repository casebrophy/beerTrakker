package controller

import (
	"net/http"
)

type User struct {
	id       int
	uName    string
	password string
}

type UserLogin struct {
	uName          string
	password       string
	sessionID      string
	loginInSuccess bool
}

func (*server) handleLogin() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

	}
}
