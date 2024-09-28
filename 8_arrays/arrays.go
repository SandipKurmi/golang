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

	// direct assign values
	var ags [4]int = [4]int{1,2,3,4}
	fmt.Println(ags)	

	var names  [4]string =  [4]string{"sandip", "uday", "sachin", "kumar"}
	fmt.Println(names)

	// 2d arrays
	var arr2d [3][3]int = [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println(arr2d)

}