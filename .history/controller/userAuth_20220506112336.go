package controller

import "net/http"

type User struct {
	id       int
	uName    string
	password string
}

func handleLogin(w http.ResponseWriter, r *http.Request) {

}
