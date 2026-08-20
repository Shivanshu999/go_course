package main

import "fmt"

func main() {
	// age := 11

	//if else statement

	//  if age >= 18 {
	// 	fmt.Println("You are eligible to vote.")
	//  } else{
	// 	fmt.Println("You are not eligible to vote.")
	//  }

	//if elseif else statement

	// if age >= 18 {
	// 	fmt.Println(" you are an adult")

	// } else if age >= 13 {
	// 	fmt.Println("you are a teenager")
	// } else {
	// 	fmt.Println("you are a child")
	// }

	// logical operators
	// we can declare variables inside the if statement
	if age := 90; age >= 18 {
		fmt.Println("you are an adult")
	} else if age > 60 {
		fmt.Println("you are a senior citizen")
	} else {
		fmt.Println("you are not an teenager")
	}

}
