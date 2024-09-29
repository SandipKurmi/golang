package main

import "fmt"

func closure() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func closure2() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {

	fn := closure()
	fmt.Println(fn()) // 1
	fmt.Println(fn()) // 2
	fmt.Println(fn()) // 3
	fmt.Println(fn()) // 4
	fmt.Println(fn()) // 5

	fn2 := closure2()
	fmt.Println(fn2()) // 1
	fmt.Println(fn2()) // 2

}