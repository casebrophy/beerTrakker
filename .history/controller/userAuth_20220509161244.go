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
	Uname          string `json:"uName"`
	Password       string `json:"password"`
	sessionID      string
	loginInSuccess bool
}

func (s *Server) handleLogin() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logging in")

		//parse request
		var u UserLogin
		b, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err := json.Unmarshal(b, &u); err != nil {
			panic(err)
		}

		fmt.Println(u.Uname, u.Password)

		//find user if they exist

		//login logic

		//set user as authenticated

	}
}

//displays the login page on get request
func (*Server) handleLoginPage() http.HandlerFunc {

	//tmpl := template.Must(template.ParseGlob("frontend/**"))

	return func(w http.ResponseWriter, r *http.Request) {

	}
}
