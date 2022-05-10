package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

type pageData struct {
	Title string
}

func (s *Server) handleHomePage() http.HandlerFunc {

	tpl := template.Must(template.ParseFiles("frontend/index.html", "templates/*"))
	fmt.Println(tpl.Templates())
	p := pageData{Title: "trakker"}
	return func(w http.ResponseWriter, r *http.Request) {
		if err := tpl.Execute(w, p); err != nil {
			fmt.Println(err.Error())
		}
	}
}
