package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) handleProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//vars is the list of variables that have been passed in to the router via GET
		vars := mux.Vars(r)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Product: %v\n", vars["product"])
	}
}
