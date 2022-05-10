package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the test server\n")
	fmt.Fprintf(w, "Request information: %v", r.Body.Close())
	fmt.Println("On homePage")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	log.Fatalln(http.ListenAndServe(":10000", nil))
}

func main() {

	fmt.Println("Starting test server")
	handleRequest()
}
