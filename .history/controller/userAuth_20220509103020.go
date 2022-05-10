package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	id       int
	uName    string
	password string
}

type UserLogin struct {
	UName          string `json:"uName"`
	Password       string `json:"password"`
	sessionID      string
	loginInSuccess bool
}

func (*server) handleLogin() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logging in")
		var u UserLogin
		b, _ := ioutil.ReadAll(r.Body)

		if err := json.Unmarshal(b, &u); err != nil {
			panic(err)
		}

		fmt.Println(u.uName, u.password)

	}
}
