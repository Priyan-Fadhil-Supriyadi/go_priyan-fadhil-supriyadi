package main

import (
	"fmt"
)

func FindMinAndMax(arr []int) string {
	// your code here
	minmax := func(numbers []int) (min int, idxmin int, max int, idxmax int) {
		// input
		min, max = numbers[0], numbers[0]
		idxmin, idxmax = 0, 0
		for i := range numbers {
			if numbers[i] > max {
				max = numbers[i]
				idxmax = i
			}
			if numbers[i] < min {
				min = numbers[i]
				idxmin = i
			}
		}
		return min, idxmin, max, idxmax
	}
	kecil, idxmin, besar, idxmax := minmax(arr)
	response := fmt.Sprintf("min: %d index: %d max: %d index: %d", kecil, idxmin, besar, idxmax)
	return response
}

func main() {
	fmt.Println(FindMinAndMax([]int{5, 7, 4, -2, -1, 8}))
	// min: -2 index: 3 max: 8 index: 5
	fmt.Println(FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))
	// min: -5 index: 1 max: 22 index: 3
	fmt.Println(FindMinAndMax([]int{4, 3, 9, 4, -21, 7}))
	// min: -21 index: 4 max: 9 index: 2
	fmt.Println(FindMinAndMax([]int{-1, 5, 6, 4, 2, 18}))
	// min: -1 index: 0 max: 18 index: 5
	fmt.Println(FindMinAndMax([]int{-2, 5, -7, 4, 7, -20}))
	// min: -20 index: 5 max: 7 index: 4
}
