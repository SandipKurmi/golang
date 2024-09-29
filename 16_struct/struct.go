package main

import (
	"fmt"
	"time"
)

// order struct
type Order struct {
	id string
	amount int
	status string
	createdAt time.Time // nanosecond precision
}

// constructor
func newOrder(id string, amount int, status string) Order {
	return Order{
		id: id,
		amount: amount,
		status: status,
		createdAt: time.Now(),
	}
}

func (o *Order) changeStatus(status string) string {
	o.status = status
	return o.status
}

// use pointer only when you want to change the value
func (o *Order) getStatus() string {
	return o.status
}

// if you are not changing the value then use value
func (o Order) getId() string {
	return o.id
}

func main() {
	// simple struct
	loginCredentials := struct {
		username string
		password string
	} {
		"sandip",
		"password",
	}

	fmt.Println("loginCredentials", loginCredentials)

	// order struct
	myOrder := Order{
		id: "1",
		amount: 100,
		status: "pending",
	}
	myOrder.createdAt = time.Now()

	// amount: 100, // if you do't set any field then it will be 
	// Int => 0 , Float => 0.0, String => "", Bool => false

	myOrder.changeStatus("paid") 

	fmt.Println(myOrder.getStatus()) 
	
	fmt.Println("order struct", myOrder)

	// accessing
	fmt.Println("id", myOrder.id)

	// if you are not changing the value then use value
	fmt.Println("order struct", myOrder.getId())

	myOrder2 := Order{
		id: "2",
		amount: 200,
		status: "paid",
	}
	myOrder2.createdAt = time.Now()
	fmt.Println("order struct", myOrder2)

	// go constructor
	myOrder3 := newOrder("3", 300, "pending")
	myOrder3.createdAt = time.Now()
	fmt.Println("order struct", myOrder3)
}