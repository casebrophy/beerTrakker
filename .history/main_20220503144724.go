package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

//structs use this interfact for data validation
type OK interface {
	ok() (err error)
}

type User struct {
	id       int
	name     string
	password string
}

func (u *User) ok() error {
	if u.id == 0 {
		return errors.New("id is required")
	}
	if u.name == "" {
		return errors.New("name is required")
	}
	if u.password == "" {
		return errors.New("password is required")
	}
	return nil
}

func decode(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return nil
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
	handleRequest()
}
