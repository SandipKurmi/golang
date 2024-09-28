package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main() {
	//creating map
	// IMP: if key does not exists in the map then it will return zero value
	city := make(map[string]string)
	city["name"] = "kolkata"
	city["population"] = "1000000"
	fmt.Println(city)

	//len
	fmt.Println(len(city))

	//delete 
	delete(city, "population")
	fmt.Println(city)

	//clear
	clear(city)
	fmt.Println(city)

	var person = map[string]string{
		"name":"sandip",
		"city":"kolkata",
	}
	person["address"] = "delhi"

	fmt.Println(person)	

	//check key exists
	if v, ok := person["name"]; ok {
		fmt.Println("key exists")
		fmt.Println(v) // sandip
	} else {
		fmt.Println("key not exists")
	}

	// check equal maps
	m1 := map[string]string{
		"name":"sandip",
		"city":"kolkata",
	}
	m2 := map[string]string{
		"name":"sandip",
		"city":"kolkata",
	}
	// fmt.Println(m1==m2) // invalid operation: m1 == m2 (map can only be compared to nil)compilerUndefinedOp
	fmt.Println(maps.Equal(m1, m2)) // true
}