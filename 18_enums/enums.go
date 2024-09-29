package main

import "fmt"

type OrderStatus string

const (
	Pending OrderStatus = "pending"
	Processing OrderStatus = "processing"
	Shipped OrderStatus = "shipped"
	Canceled OrderStatus = "canceled"
)

type Role string

const (
	User Role = "user"
	Admin Role = "admin"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println(status)
}

func isAdmin(role Role) bool {
	return role == Admin
}

func main() {
	changeOrderStatus(Processing)	

	if isAdmin(Admin) {
		fmt.Println("admin") // admin
	}
}