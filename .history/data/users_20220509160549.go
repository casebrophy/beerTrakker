package data

import (
	"testServer/controller"
)

func getUser(username string) (*controller.User, error) {
	
	if err := db.QueryRow("SELECT username, password FROM users WHERE username = ?", username); err != nil {
		if err
	}
}