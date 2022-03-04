package main

import (
	"fmt"
	"math"
)

func primaSegiEmpat(high, wide, start int) {
	// your code here
	var total, i, j int
	start += 1
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

	for i < wide {
		if prime(start) {
			j++
			total += start
			fmt.Print(start)
			if start <= 9 {
				fmt.Print(" ")
			} else if start > 99 {
				fmt.Print(" ")
				continue
			}
			fmt.Print("  ")
		}
		if j == high {
			i++
			j = 0
			fmt.Println()
		}
		start += 1
	}
	fmt.Println(total)
	fmt.Println()
}

func main() {
	primaSegiEmpat(2, 3, 13)
	/*
	   17 19
	   23 29
	   31 37
	   156
	*/
	primaSegiEmpat(5, 2, 1)
	/*
	   2  3  5  7 11
	   13 17 19 23 29
	   129

	*/

}
