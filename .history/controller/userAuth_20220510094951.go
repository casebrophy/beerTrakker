package controller

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	id           int
	uName        string
	passwordHash string
}

type UserLogin struct {
	Uname          string `json:"uName"`
	Password       string `json:"password"`
	PasswordHash   [32]byte
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
			fmt.Println(err)
		}

		fmt.Println(u.Uname, u.Password)

		u.PasswordHash = sha256.Sum256([]byte(u.Password))

		u.Password = "" //clear password to limit time in plaintext

		//find user if they exist
		user, err := s.GetUser(u.Uname)

		if err != nil {
			fmt.Println(err)
		}

		//login logic
		status := checkPassword(user, u.Password)

		//set user as authenticated

	}
}

//displays the login page on get request
func (*Server) handleLoginPage() http.HandlerFunc {

	//tmpl := template.Must(template.ParseGlob("frontend/**"))

	return func(w http.ResponseWriter, r *http.Request) {

	}
}
