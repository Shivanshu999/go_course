package main

import "fmt"

func calculate(a float64, b float64, opreator string) float64 {

	switch opreator {
	case "+":
		return a + b

	case "-":
		return a - b

	case "/":
		return a / b

	case "*":
		return a * b

	default:
		fmt.Println("Invalid operator")
		return 0
	}
}

func main() {
	var num1 float64
	var num2 float64
	var opreator string

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter Operator (+, - , * , /)")
	fmt.Scan(&opreator)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	result := calculate(num1, num2, opreator)
	fmt.Println("Result:", result)
}
