package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the test server")
	fmt.Println("On homePage")
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {

	fmt.Println("Starting test server")

}
