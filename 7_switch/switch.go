package main

import (
	"fmt"
	"time"
)

func main(){
	i := 2

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Other")
	}

	//multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's workday",time.Now().Weekday())
	}

	//type switch

	whoAmI := func(i interface{})  {
		switch t := i.(type){
		case int:
			fmt.Println("it's an integer")
		case string:
			fmt.Println("it's an string")
		case bool:
			fmt.Println("it's an boolean")
		default:
			fmt.Println("other", t)
		}
	}

	whoAmI(false)
	whoAmI("Sandip")
	whoAmI(07)
}