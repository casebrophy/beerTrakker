package controller

import (
	"fmt"
)

//checks for user in database
//PARAMS username
//RETURNS User struct
func (s *Server) GetUser(username string) (*User, error) {

	result := s.db.QueryRow("SELECT username, password FROM users WHERE username = ?", username)
	fmt.Println(result)
	return nil, nil

}
