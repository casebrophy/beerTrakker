package controller

import (
	"fmt"
	"net/http"
	"text/template"
)

type data struct {
	title      string
	Controller string
}

func (s *server) handleHomePage() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("templates/*"))
	data := data{title: "test", Controller: "test"}
	fmt.Println("in template")
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	}
}
