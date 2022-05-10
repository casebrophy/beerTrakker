package data

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

//called by main to initializeDB
func initializeDB() *sql.DB {

	db := &sql.DB{}
	var err error

	cfg := mysql.Config{
		User:                 os.Getenv("dbUSER"),
		Passwd:               os.Getenv("dbPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               os.Getenv("dbNAME"),
		AllowNativePasswords: true,
	}

	if db, err = sql.Open("mysql", cfg.FormatDSN()); err != nil {
		log.Fatal(err)
	}

	return db
}
