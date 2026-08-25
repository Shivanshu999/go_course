// Asignment:  use Go’s crypto/rand package should be used because it provides cryptographically secure randomness.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func passwordGenerator(length int) string {
	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
	password := ""
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < length; i++ {
		randomIndex := r.Intn(len(char))
		password += string(char[randomIndex])

	}
	return password
}

func main() {
	var length int

	fmt.Print("enter the password length you want: ")
	fmt.Scan(&length)
	if length <= 0 {
		fmt.Println("Password length should be greater then zero")
		return
	}

	password := passwordGenerator(length)
	fmt.Println("your password is: ", password)
}
