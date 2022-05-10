package controller

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
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
		status := CheckPassword(user, u.Password)

		if status {
			fmt.Println("Success")
		} else {
			fmt.Println("Can't login")
		}

		//set user as authenticated

	}
}

//displays the login page on get request
func (*Server) handleLoginPage() http.HandlerFunc {

	t := GetTemplates(".frontend/login.http")

	tmpl := template.Must(template.ParseFiles(t...))

	return func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.Execute(w, r); err != nil {
			panic(err)
		}
	}
}
