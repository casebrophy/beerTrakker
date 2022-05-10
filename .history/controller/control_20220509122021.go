package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

type pageData struct {
	Title string
}

func (s *server) handleHomePage() http.HandlerFunc {

	var tpl = template.Must(template.ParseFiles("frontend/index.html", "frontend/menu.html", "frontend/footer.html", "frontend/header.html"))
	p := pageData{Title: "trakker"}
	return func(w http.ResponseWriter, r *http.Request) {
		if err := tpl.Execute(w, p); err != nil {
			fmt.Println(err.Error())
		}
	}
}
