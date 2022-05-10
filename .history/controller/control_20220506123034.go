package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *server) handleHomePage(w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	data := "Hello world!"
	fmt.Fprintf(w, "%s\n", data)
	vars := mux.Vars(r)

	for _, k := range vars {
		fmt.Println("Key: ", k, "Value: ", vars[k])
	}
}
