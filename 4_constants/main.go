package main

import "fmt"

const age = 24
const name = "sandip"
func main(){
	const name = "sandip"
	const age = 30
	
	// age=24 //we can not assign 
	// name = "uday" //we can not assign 

	fmt.Println(age)
	fmt.Println(name)

	const (
		port = 3000
		host = "localhost"
	)

	fmt.Println(port)
	fmt.Println(host)
}