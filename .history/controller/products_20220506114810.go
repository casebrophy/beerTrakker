package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Product: %v\n", vars["product"])
}
