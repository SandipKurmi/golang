package main

import (
	"fmt"
)

func main() {
	// slice
	nums := []int{1, 2, 3}
	sum := 0
	for i, num := range nums {
		sum += num
		fmt.Println(i,num)
	}
	fmt.Println(sum)

	//maps
	m := map[string]string{
		"fristName":"sandip",
		"lastName":"kurmi",
	}

	for k, v := range m {
		fmt.Println(k,v)
	}

	// string range
	// unicode code point rune
	// starting byte of rune
	// 255 -> 1 byte
	// 0 -> 1 byte
	s := "sandip"
	for i, v := range s {
		fmt.Println(i,v, string(v))
		// 0 115 s
		// 1 97 a
		// 2 110 n
		// 3 100 d
		// 4 105 i
		// 5 112 p
	}

}

