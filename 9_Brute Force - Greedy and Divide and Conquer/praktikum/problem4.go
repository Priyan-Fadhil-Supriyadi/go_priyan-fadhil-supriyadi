package main

import "fmt"

func BinarySearch(array []int, x int) {
	// your code here
	var min, max = 0, len(array) - 1
	for min <= max {
		mid := (min + max) / 2
		if array[mid] == x {
			fmt.Println(mid)
			return
		} else if x < array[mid] {
			max = mid
		} else {
			min = mid + 1
		}
	}
	fmt.Println(-1)
}

func main() {
	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3)                  // 2
	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5)                 // 3
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53)  // 6
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100) // -1
}
