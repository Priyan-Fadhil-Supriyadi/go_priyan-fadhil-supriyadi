package main

import (
	"fmt"
	"math"
)

func primeNumber(number int) bool {

	log := math.Sqrt(float64(number))
	for i := 2.0; i < log; i++ {
		n := int(i)
		if number%n == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(13))         // true
	fmt.Println(primeNumber(17))         // true
	fmt.Println(primeNumber(20))         // false
	fmt.Println(primeNumber(35))         // false
}
