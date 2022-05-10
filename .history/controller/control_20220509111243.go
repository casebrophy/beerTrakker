package controller

import (
	"html/template"
	"net/http"
)

func (s *server) handleHomePage() http.HandlerFunc {

	var tpl = template.Must(template.ParseFiles("frontend/index.html", "frontend/menu.html", "frontend/footer.html", "frontend/header.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
