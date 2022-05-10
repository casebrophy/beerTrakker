package controller

import (
	"net/http"
	"text/template"
)

func (s *server) handleHomePage() http.HandlerFunc {

	var tpl = template.Must(template.ParseFiles("frontend/index.html", "frontend/menu.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
