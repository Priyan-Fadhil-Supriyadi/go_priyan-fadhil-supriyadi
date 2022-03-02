package main

import (
	"fmt"
)

func getMinMax(numbers ...*int) (min int, max int) {
	// input
	min, max = *numbers[0], *numbers[0]
	for i := range numbers {
		if *numbers[i] > max {
			max = *numbers[i]
		}
		if *numbers[i] < min {
			min = *numbers[i]
		}
	}
	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai min ", min)
	fmt.Println("Nilai max ", max)
}
