package main

import "fmt"

type OrderStatus string

const (
	Recived    OrderStatus = "Recived"
	Processing OrderStatus = "Processing"
	Shipped    OrderStatus = "Shipped"
	Delivered  OrderStatus = "Delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order status changed to: ", status)
}

func main() {
	changeOrderStatus(Recived)
	changeOrderStatus(Processing)
	changeOrderStatus(Shipped)
	changeOrderStatus(Delivered)
}
