package main

import "fmt"

// type : 1
// func add(a int, b int) int {
// 	return a + b
// }

// we can write above function as this also
// func add(a, b int) int {
// 	return a + b
// }

// func main() {
// 	result := add(2, 1)
// 	fmt.Println(result)
// }

// type : 2

// func getLanguage() (string, string, bool) {
// 	return "golang", "python", true
// }

// func main() {
// 	lang1, lang2, lang3 := getLanguage()
// 	fmt.Println(lang1, lang2, lang3)

// }

// type : 3

// func processIt(fn func(a int) int) {
// 	fn(1)
// }

// func main() {
// 	fn := func(int) int {
// 		return 2
// 	}
// 	fmt.Println(fn(1))
// }

//type : 4

func doSomething() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {
	fn := doSomething()
	fn(6)

	fmt.Println(fn(1))
}
