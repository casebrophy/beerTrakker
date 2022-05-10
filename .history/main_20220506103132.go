package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

//structs use this interfact for data validation
type OK interface {
	OK() error
	Error() string //not sure why this is needed but I cannot check if something i
}

func homePage(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respond(w, r, http.StatusInternalServerError, err)
		return
	}
	respond(w, r, 300, data)
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	log.Fatalln(http.ListenAndServe(":10000", nil))
}

func respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	if _, err := io.Copy(w, &buf); err != nil {
		log.Println("respond:", err)
	}

}

func main() {

	fmt.Println("Starting test server")

	//loading in environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Env variables USER: %s\n", os.Getenv("USER"))
	//initalize db

	//handleRequest()
}
