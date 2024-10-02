package auth

import "fmt"

type Credentials struct {
	Username string
	Password string
}

// in go if you start with capital letter then it is public
// if you start with small letter then it is private

func NewCredentials(username, password string) *Credentials {
	fmt.Println("NewCredentials")
	return &Credentials{
		Username: username,
		Password: password,
	}
}