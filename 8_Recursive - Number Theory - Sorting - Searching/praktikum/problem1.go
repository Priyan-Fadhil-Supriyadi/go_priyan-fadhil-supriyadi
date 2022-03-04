package main

import (
	"fmt"
	"math"
)

func primeX(number int) int {
	if number == 1 {
		return 2
	}
	cek, i, next := 2, 0, 0
	prime := func(cek int) bool {
		log := int(math.Sqrt(float64(cek)))
		for i := 2; i <= log; i++ {
			n := int(i)
			if cek%n == 0 {
				return false
			}
		}
		return true
	}
	for i < number {
		pass := prime(cek)
		if pass == true {
			next = cek
			i++
		}
		cek++
	}
	return next
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29

}
