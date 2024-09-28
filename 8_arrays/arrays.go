package main

import "fmt"

func main() {
	var nums [4]int

	//int -> 0 
	nums[0] = 1
	fmt.Println(nums)
	fmt.Println(nums[0])
	
	//array length
	fmt.Println(len(nums))
	

	// bool array
	var vals [4]bool
	vals[2] = true
	fmt.Println(vals)

	// string array
	var valsstr [4]string
	valsstr[2] = "sandip"
	fmt.Println(valsstr)
}