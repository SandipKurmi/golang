package main

import "fmt"


func main(){
	age:=22

	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else {
	// 	fmt.Println("is not an adult")
	// }

	if age >= 18 {
		fmt.Println("Person is an adult")
	} else if age >= 12 {
		fmt.Println("Person is teenager")
	} else {
		fmt.Println("Person is a kid")
	}


	var role = "admin"

	var hasPermission = false

	if role == "admin" && hasPermission {
		fmt.Println("Yes")
	}


	//we can declare a variable inside if construct.
	if age:=10 ; age >= 18 {
		fmt.Println("Persion is and adult", age)
	} else if age >= 12 {
		fmt.Println("Person is teenager", age)
	} else {
		fmt.Println("Person is kid", age)
	}

// go does not have ternary, you will hvae to use normal if else

}