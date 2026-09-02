package main

import (
	"fmt"
)

// send
// func processNum(numChan chan int) {

// 	for num := range numChan {
// 		fmt.Println("Processing number", num)
// 		time.Sleep(time.Second)

// 	}
// 	// numChan := make(chan int)
// 	// go processNum()

// }
// func main() {

// 	numChan := make(chan int)
// 	go processNum(numChan)

// 	for {
// 		numChan <- rand.Intn(100)
// 	}

// 	// messageChan := make(chan string)

// 	// messageChan <- "ping"

// 	// msg := <-messageChan
// 	// fmt.Println(msg)
// }

// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2

// 	result <- numResult
// }

// func main() {
// 	result := make(chan int)
// 	go sum(result, 4, 5)
// 	res := <-result
// 	fmt.Println(res)

// }

// goroutine synchronizer

// func task(done chan bool) {
// 	defer func() { done <- true }()
// 	fmt.Println("process...")
// }

// func main() {
// 	done := make(chan bool)
// 	go task(done)

// 	<-done
// }

//that's how we can implement queu system in golang
// func emailSender(emailChan chan string, done chan bool) {
// 	defer func() { done <- true }()
// 	for email := range emailChan {
// 		fmt.Println("sending email to", email)
// 		time.Sleep(time.Second)
// 	}
// }
// func main() {
// 	emailChan := make(chan string, 100)
// 	done := make(chan bool)
// 	go emailSender(emailChan, done)
// 	for i := 0; i < 100; i++ {
// 		emailChan <- fmt.Sprintf("%d@gmail.com", i)
// 	}
// 	fmt.Println(("done"))
// 	close(emailChan)
// 	<-done

// 	// emailChan <- "1@exmple.com"
// 	// emailChan <- "2@exmple.com"
// 	// fmt.Println(<-emailChan)
// 	// fmt.Println(<-emailChan)
// }

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10

	}()
	go func() {
		chan2 <- "ping"

	}()
	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("recived data from chan1", chan1Val)

		case chan2Val := <-chan2:
			fmt.Println("recived data from chan2", chan2Val)
		}
	}

}
