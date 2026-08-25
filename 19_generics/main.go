package main

import "fmt"

// func printSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }
// func printString(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func main() {
// 	// printSlice([]int{1, 2, 3})
//     //printString([]string{"go", "js"})
// 	nums := []int{1, 2, 3}
// 	names := []string{"golang", "ts"}
// 	printString(names)
// 	printSlice(nums)
// }

// the code above is not good practice the code below will is better with the help of using generics
// when you wanna use slice
func printSlice[T string | int | bool](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// when you wanna use struct
type stack[T any] struct {
	elements []T
}

func main() {

	myStack := stack[string]{
		elements: []string{"golang", "rust", "C#"},
	}

	// printSlice([]int{1, 2, 3})
	//printString([]string{"go", "js"})
	nums := []int{1, 2, 3}
	names := []string{"golang", "ts"}
	boolean := []bool{true, false, true}
	printSlice(names)
	printSlice(nums)
	printSlice(boolean)
	fmt.Println(myStack)
}
