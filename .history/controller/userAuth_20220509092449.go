package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	id       int
	uName    string
	password string
}

func (*server) handleLogin() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/home.html"))
		data := data{Controller: "test"}
		fmt.Println("in template")
		tmpl.Execute(w, data)
		//fmt.Fprintln(w, "You're home")
	}
}
