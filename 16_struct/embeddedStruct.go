package main

import "fmt"

type User struct {
	name string
	age int
}


// embedded struct
type Cart struct {
	User
	products []string
}

func cartMain() {

	// embedded struct
	cart := Cart{
		User: User{
			name: "sandip",
			age: 24,
		},
		products: []string{"mobile", "laptop"},
	}
	fmt.Println(cart)
}