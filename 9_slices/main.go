package main

import (
	"fmt"
	"slices"
)

// slice -> dynimic array
// most used construct in go
// + useful method

func main() {
	//uninitialized slice is nil
	var nums []int 
	fmt.Println(nums) //[]
	fmt.Println(len(nums)) //0
	nums = append(nums, 1)
	fmt.Println(nums)//[1]
	fmt.Println(len(nums)) //1
	fmt.Println(nums == nil) //false

	var age = make([]int, 0, 4)
	fmt.Println(age) 
	age = append(age, 1)
	age = append(age, 2)
	age = append(age, 3)
	age = append(age, 4)
	fmt.Println(age) // [1,2,3,4]

	// capacity => maximum numbers of element can fit
	fmt.Println(cap(age)) // dynimic capacity

	// copy function
	src := []int{1, 2, 3, 4, 5}
	dest := make([]int, len(src))
	copy(dest, src)

	fmt.Println(src) // [1 2 3 4 5]
	fmt.Println(dest) // [1 2 3 4 5]

	// slicing
	var name []string = []string{"sandip", "uday", "sachin", "kumar"}
	fmt.Println(name[0:3]) // [sandip uday]

	// append	
	name = append(name, "kumar")

	// slices package
	var nums1 []int = []int{1, 2, 3, 4, 5}
	var nums2 []int = []int{6, 7, 8, 9, 10}

	fmt.Println(slices.Equal(nums1, nums2))

	// 2d slices
	var arr [][]int = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(arr)
}