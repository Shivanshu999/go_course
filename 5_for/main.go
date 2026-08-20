package main

import "fmt"

func main() {
	// while loop

	// i := 0
	// for i < 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	//infinite loop

	// for {
	// 	fmt.Println("infinite loop")
	// }

	//

	//classic for loop

	// for i := 0; i < 5; i++ {
	// 	//     break

	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	//go 1.22 update range

	for i := range 8 {
		fmt.Println(i)
	}

}
