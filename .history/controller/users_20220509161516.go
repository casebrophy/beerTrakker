package controller

import (
	"fmt"
)

func (s *Server) getUser(username string) (*User, error) {

	result := s.db.QueryRow("SELECT username, password FROM users WHERE username = ?", username)
	fmt.Println(result)
	return nil, nil

}
