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
}

func (*server) handleLogin() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseGlob("./templates/*"))
		d := data{Content: "test"}
		fmt.Println(tmpl)
		tmpl.Execute(w, d)
		//fmt.Fprintln(w, "You're home")
	}
}
