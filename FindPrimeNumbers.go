package main

import (
	"fmt"
	"math"
)

// Function to check if numbers are prime or non-prime numbers
func checkPrime(num int) {
	sq_root := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq_root; i++ {
		// check if its a non prime number
		if num%i == 0 {
			return
		}
	}
	// else if its a prime return the number
	fmt.Println(num)
	return
}

func main() {
	for i := 2; i <= 100; i++ {
		checkPrime(i)
	}
}
