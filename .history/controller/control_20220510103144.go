package controller

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type pageData struct {
	Title string
}

func (s *Server) handleHomePage() http.HandlerFunc {

	t := GetTemplates()

	tpl := template.Must(template.ParseFiles("frontend/index.html", t))
	fmt.Println(tpl.Templates())
	p := pageData{Title: "trakker"}
	return func(w http.ResponseWriter, r *http.Request) {
		if err := tpl.Execute(w, p); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func GetTemplates() []string {
	var templates []string
	files, err := ioutil.ReadDir("./templates")
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, file := range files {
		filename := file.Name()
		templates = append(templates, "./templates/"+filename)
	}
	return templates
}
