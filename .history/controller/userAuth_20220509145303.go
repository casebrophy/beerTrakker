package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/sessions"
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

func (s *server) handleLogin() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logging in")
		var u UserLogin
		b, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		if err := json.Unmarshal(b, &u); err != nil {
			panic(err)
		}

		fmt.Println(u.Uname, u.Password)
		var session *sessions.Session
		var err error
		if session, err = s.store.Get(r, u.Uname); err != nil {
			panic(err)
		}

	}
}

//displays the login page on get request
func (*server) handleLoginPage() http.HandlerFunc {

	//tmpl := template.Must(template.ParseGlob("frontend/**"))

	return func(w http.ResponseWriter, r *http.Request) {

	}
}
