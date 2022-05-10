package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
	data := "Hello world!"
	fmt.Fprintf(w, "%s\n", data)
}
