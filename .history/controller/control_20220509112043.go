package controller

import (
	"html/template"
	"net/http"
)

type pageData struct {
	title string
}

func (s *server) handleHomePage() http.HandlerFunc {

	var tpl = template.Must(template.ParseFiles("frontend/index.html", "frontend/menu.html", "frontend/footer.html", "frontend/header.html"))
	p := pageData{title: "trakker"}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
