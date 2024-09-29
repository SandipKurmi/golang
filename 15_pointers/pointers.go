package main

import "fmt"

// pass by value
func changeByValue(num int ) {
	num = 10
	fmt.Println("Value inside function", num)
}

// pass by reference * is a pointer
func changeByReference(num *int ) {
	*num = 10
	fmt.Println("Value inside function", *num)
}

func changeByReferenceSlice(nums *[]int) {
	*nums = append(*nums, 10)
	fmt.Println("Value inside function", *nums) // [1 2 3 4 5 10]
}

func changeByReferenceNameSlice(names *[]string) {
	*names = append(*names, "deepak")
	fmt.Println("Value inside function", *names) // [sandip uday sachin kumar sandip]
}



func main() {
	num := 5

	// pass by value
	changeByValue(num)

	fmt.Println("Value outside function", num)
	// pass by reference
	changeByReference(&num) // &num is a pointer

	fmt.Println("Value outside function", num)

	// number slice
	nums := []int{1,2,3,4,5}

	fmt.Println(nums) // [1 2 3 4 5]

	changeByReferenceSlice(&nums)

	fmt.Println("Value outside function", nums) // [1 2 3 4 5 10]

	// string slice
	names := []string{"sandip", "uday", "sachin", "kumar"}

	fmt.Println(names) // [sandip uday sachin kumar]

	changeByReferenceNameSlice(&names)

	fmt.Println("Value outside function", names) // [sandip uday sachin kumar deepak]
}