package main

import (
	"fmt"
	"time"
)

// func task(id int) {
// 	fmt.Println("Doing task", id)

// }

func main() {
	for i := 0; i <= 10; i++ {
		//you can write code below
		// go task(i)
		// or like this
		go func(i int) {
			fmt.Println(i)
		}(i)

	}
	time.Sleep(time.Second * 2)
}
