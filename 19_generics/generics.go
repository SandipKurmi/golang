package main

import "fmt"

// not generics
func printNumberList(nums []int){
	for _, v := range nums {
		fmt.Println(v)
	}
}

func printStringList(names []string){
	for _, v := range names {
		fmt.Println(v)
	}
}

func printList[T int | string | bool](nums []T){
	for _, v := range nums {
		fmt.Println(v)
	}
}

type Stack[T any] struct {
	elements []T
}

func main() {
	// int slice
	nums := []int{1,2,3,4,5}
	printNumberList(nums)
	printList(nums)

	// string slice
	names := []string{"sandip", "uday", "sachin", "kumar"}
	printStringList(names)
	printList(names)

	// bool slice
	bools := []bool{true, false, true}
	printList(bools)

	// generics stack
	myStackNumber := Stack[int]{
		elements: []int{1,2,3,4,5},
	}

	fmt.Println(myStackNumber.elements)

	myStackString := Stack[string]{
		elements: []string{"sandip", "uday", "sachin", "kumar"},
	}

	fmt.Println(myStackString.elements)
}