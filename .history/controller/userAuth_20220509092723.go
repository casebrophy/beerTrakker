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

type data struct {
	Controller string
}

func (*server) handleLogin() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/home.html"))
		d := data{Controller: "test"}
		fmt.Println("in template")
		tmpl.Execute(w, d)
		//fmt.Fprintln(w, "You're home")
	}
}
