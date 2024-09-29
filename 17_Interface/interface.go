package main

import "fmt"

// interface is a contract
type paymentGateway interface {
	pay(amount int)
	refund(amount int)
}

type payment struct{
	gateway paymentGateway
}

func (p payment) makePayment(amount int) {
	p.gateway.pay(amount)
}


func (p payment) refund(amount int) {
	p.gateway.refund(amount)
}

type razorpay struct {}

func (r razorpay) pay(amount int) {
	fmt.Println(" Make payment using Razorpay", amount)
}

type stripe struct {}

func (s stripe) pay(amount int) {
	fmt.Println(" Make payment using Stripe", amount)
}

type fackPaymentGateway struct {}

func (f fackPaymentGateway) pay(amount int) {
	fmt.Println(" Make payment using fackPaymentGateway", amount)
}

type paypal struct {}

func (p paypal) pay(amount int) {
	fmt.Println(" Make payment using paypal", amount)
}

func (p paypal) refund(amount int) {
	fmt.Println(" Make refund using paypal", amount)
}

func main() {

	// stripePayment := stripe{}
	// razorpayPayment := razorpay{}
	// fackPaymentGateway := fackPaymentGateway{}
	paypalPayment := paypal{}
	payment := payment{
		gateway: paypalPayment,
	}
	
	payment.makePayment(100)
	payment.refund(50)
}