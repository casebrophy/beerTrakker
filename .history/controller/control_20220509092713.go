package controller

import (
	"net/http"
)

func (s *server) handleHomePage() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bruh"))
	}
}
