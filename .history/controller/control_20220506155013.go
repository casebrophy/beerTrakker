package controller

import (
	"net/http"
	"text/template"
)

type data struct {
	title      string
	Controller string
}

func (s *server) handleHomePage() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	data := data{title: "test", Controller: "test"}
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	}
}
