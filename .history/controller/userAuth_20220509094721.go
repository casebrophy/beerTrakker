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
	Content string
	title   string
}

func (*server) handleLogin() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseGlob("./templates/*"))
		d := data{Content: "test", title: "title"}
		fmt.Println(&tmpl)
		tmpl.Execute(w, d)
		//fmt.Fprintln(w, "You're home")
	}
}
