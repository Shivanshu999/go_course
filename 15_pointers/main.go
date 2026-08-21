package main

import "fmt"

// by value

// func changeValue(x int) {
// 	x = 10
// 	fmt.Println("Inside changeValue function:", x)
// }

// func main() {
// 	x := 5

// 	changeValue(x)
// 	fmt.Println("After changeValue function:", x)
// }

// by reference => here is a pointer to the variable x, so we can change the value of x in the changeValue function
func changeValue(x *int) {
	*x = 10
	fmt.Println("Inside changeValue function:", *x)
}

func main() {
	x := 5
	changeValue(&x)
	fmt.Println("After changeValue function:", x)
}
