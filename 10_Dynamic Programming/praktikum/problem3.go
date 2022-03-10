package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	// your code here
	var arr = make(map[int]int)
	arr[1] = int(math.Abs(float64(jumps[1] - jumps[0])))
	for i := 2; i < len(jumps); i++ {
		arr[i] = int(math.Min(math.Abs(float64(jumps[i]-jumps[i-1]))+float64(arr[i-1]),
			math.Abs(float64(jumps[i]-jumps[i-2]))+float64(arr[i-2])))
		//fmt.Println("isi arr : ", arr[i-1], arr[i])
	}
	//fmt.Println(len(jumps) - 1)
	return arr[len(jumps)-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
