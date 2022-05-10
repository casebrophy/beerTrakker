package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"testServer/controller"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

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

	//initalize db
	cfg := mysql.Config{
		User:                 os.Getenv("dbUSER"),
		Passwd:               os.Getenv("dbPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               os.Getenv("dbNAME"),
		AllowNativePasswords: true,
	}

	var db *sql.DB
	var err error

	if db, err = sql.Open("mysql", cfg.FormatDSN()); err != nil {
		log.Fatal(err)
	}

	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connection established")

	controller.Routes()
}
