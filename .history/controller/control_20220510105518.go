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

	t := GetTemplates("./frontend/index.html")

	tpl := template.Must(template.ParseFiles(t...))
	fmt.Println(t)
	fmt.Println(tpl.DefinedTemplates())
	p := pageData{Title: "trakker"}
	return func(w http.ResponseWriter, r *http.Request) {
		if err := tpl.Execute(w, p); err != nil {
			fmt.Println(err)
		}
	}
}

//gets all templates from the tempalates directory
//and appends the page you would like to serve to the
//user
//PARAMETERS: path of page to serve
//RETURNS: array of strings
func GetTemplates(page string) []string {
	var templates []string
	files, err := ioutil.ReadDir("./templates")
	if err != nil {
		fmt.Println(err.Error())
	}
	//top level page must be first
	templates = append(templates, page)
	for _, file := range files {
		filename := file.Name()
		templates = append(templates, "./templates/"+filename)
	}

	return templates
}
