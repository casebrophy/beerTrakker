package controller

import (
	"net/http"
)

type data struct {
	title      string
	Controller string
}

func (s *server) handleHomePage() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

	}
}
