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
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	data := data{title: "test", Controller: "test"}
	fmt.Println("in template")
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
		//fmt.Fprintln(w, "You're home")
	}
}
