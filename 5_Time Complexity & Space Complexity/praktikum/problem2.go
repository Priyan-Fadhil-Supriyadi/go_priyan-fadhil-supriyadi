package main

import (
	"fmt"
	"math"
)

func pow(x, n int) int {
	var exp float64
	exp = math.Pow((float64(x)), float64(n))

	return int(exp)
}

func main() {
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
}
