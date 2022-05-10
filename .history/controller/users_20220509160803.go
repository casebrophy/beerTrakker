package data

import (
	"database/sql"
	"fmt"
	"testServer/controller"
)

func getUser(username string) (*controller.User, error) {

	if result := db.QueryRow("SELECT username, password FROM users WHERE username = ?", username); err != nil {
		if result == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
	} else {
		fmt.Println(result)
	}
}
