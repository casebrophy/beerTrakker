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
	tmpl := template.Must(template.ParseFiles("templates/home.html", "templates/footer.html", "templates/header.html", "templates/menu.html"))
	data := data{title: "test", Controller: "test"}
	fmt.Println("in template")
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	}
}
