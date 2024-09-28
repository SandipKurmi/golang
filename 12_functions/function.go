package main

import "fmt"

//in go function are frist class citizen

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mult(x, y int) int {
	return x * y
}

func div(x, y int) int {
	return x / y
}

func getLanguage() (string, string, bool) {
	return "Go", "JavaScript", true
}

func processIt(fn func(int, int) int, x, y int) int {
	return fn(x, y)
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(sub(1, 2))
	fmt.Println(mult(1, 2))
	fmt.Println(div(1, 2))
	fmt.Println(getLanguage())
	lan1, lan2, lan3 := getLanguage()
	fmt.Println(lan1, lan2, lan3)

	//function as parameter
	var fn func(int, int) int
	fn = add
	fmt.Println(processIt(fn, 1, 2))

	//function as parameter with anonymous function
	fmt.Println(processIt(func(x, y int) int {
		return x * y
	}, 2, 2))
}