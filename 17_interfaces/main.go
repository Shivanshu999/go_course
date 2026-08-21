package main

import "fmt"

type payment struct{}

func (p payment) processPayment(amount float32) {
	razorpayPaymentGw := razorpay{}
	razorpayPaymentGw.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	//logic to make payment

	fmt.Println("making payment using razorpay for amount: ", amount)
}

func main() {
	newPaymnet := payment{}
	newPaymnet.processPayment(100)
}
