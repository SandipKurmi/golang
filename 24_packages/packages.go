package main

import (
	"SandipKurmi/golang/24_packages/auth"
	"SandipKurmi/golang/24_packages/user"
	"fmt"

	"github.com/fatih/color"
)


func main() {
	credentials := auth.NewCredentials("sandip", "sandip")

	fmt.Println(credentials.Password)


	session := auth.GetSession("sandip", "sandip")

	fmt.Println(session)
	fmt.Println(session.Email)

	// user package

	user := user.NewUser("sandip", 24)

	fmt.Println(user)

	color.Red("Prints text in cyan.") // external package
}