package data

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

//called by main to initializeDB
func initializeDB() *sql.DB {

	cfg := mysql.Config{
		User:                 os.Getenv("dbUSER"),
		Passwd:               os.Getenv("dbPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               os.Getenv("dbNAME"),
		AllowNativePasswords: true,
	}
}
