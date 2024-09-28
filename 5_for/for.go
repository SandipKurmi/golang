package main

import "fmt"

// for -> only construct in go for looping
func main(){
	//while loop
	// j := 1
	// for j <= 3 {
	// 	fmt.Println(j)
	// 	j = j + 1
	// }

	//infinite loop

	// for {
	// 	println('1')
	// }

	// classic for loop

	// for i := 0; i <= 3 ; i++ {
	// 	fmt.Println(i)
	// }

	// classic for loop with brack

	// for i := 0; i <= 3 ; i++ {
	// 	fmt.Println(i)
	// 	break
	// }

	// classic for loop with continue

	// const count = 10

	// for i := 0; i < count; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	//10 is inclusive
	for i:= range 10 +1  {
		fmt.Println(i)
	}

}