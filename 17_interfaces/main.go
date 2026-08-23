package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}
type payment struct {
	gatway paymenter
}

func (p payment) processPayment(amount float32) {
	p.gatway.pay(amount)
}

// type razorpay struct{}

// func (r razorpay) pay(amount float32) {
// 	//logic to make payment

// 	fmt.Println("making payment using razorpay for amount: ", amount)
// }

// type stripe struct{}
// func (s stripe) pay(amount float32) {
// 	//logic to make payment

// 	fmt.Println("making payment using stripe for amount: ", amount)
// }

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal for amount: ", amount)
}

//for testing purpose we are not using any payment gateway

// type testPayment struct{}

// func (t testPayment) pay(amount float32) {
// 	fmt.Println("making payment using testPayment for amount: ", amount)
// }

func main() {
	// razorpayPaymentGw := razorpay{}
	// stripePaymentGw := stripe{}
	// testPayment := testPayment{}
	paypalPayment := paypal{}
	newPaymnet := payment{
		// gatway: razorpayPaymentGw,
		// gatway: stripePaymentGw,
		// gatway: testPayment,
		gatway: paypalPayment,
	}

	newPaymnet.processPayment(100)
}
