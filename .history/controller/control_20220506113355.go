package controller

import (
	"fmt"
	"net/http"
)

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	data := "Hello world!"
	fmt.Fprintf(w, "%s\n", data)
}
