package main

import "fmt"

const wrestlerName string = "John Cena"
const superMove = "Attitude Adjustment"

// submissionMove := "STF" => this shorthand syntax is not allowed outside of a function
var newWrestlerName string = "The Rock"
var newWrestlerMove = "Rock Bottom"

// submissionMove := "STF" // this shorthand syntax is allowed inside of a function

func main() {
	const name = "golang"
	school := "golang school"
	const isAdult = true

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(name)
	fmt.Println(school)
	fmt.Println(isAdult)
	fmt.Println(wrestlerName)

	fmt.Println(port, host)
}
