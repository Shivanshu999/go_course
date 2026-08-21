package main

import "fmt"

func doSomething(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total

}

// func main() {
// 	result := doSomething(1, 2, 3, 4, 5)
// 	fmt.Println(result)
// }

func main() {
	nums := []int{1, 2, 3, 4, 5}
	result := doSomething(nums...)
	fmt.Println(result)
}
