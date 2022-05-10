package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	id       int
	uName    string
	password string
}

type UserLogin struct {
	uName          string `json:"uName"`
	password       string `json:"password"`
	sessionID      string
	loginInSuccess bool
}

func (*server) handleLogin() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logging in"")
		decoder := json.NewDecoder(r.Body)
		var u UserLogin
		if err := decoder.Decode(&u); err != nil {
			panic(err)
		}

		fmt.Println(u.uName, u.password)

	}
}
