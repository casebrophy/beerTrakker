package testserver

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting test server")

	type Handler interface {
		ServeHTTP(ResponseWriter, *Request)
	}

	type HandlerFunc func(ResponseWriter, *Request)
	
	func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *Request) {
		f(w, r)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write()
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}