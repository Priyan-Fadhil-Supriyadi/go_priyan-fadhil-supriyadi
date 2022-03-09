package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	// your code here
	var dp = make(map[int]int)
	dp[1] = int(math.Abs(float64(jumps[1] - jumps[0])))
	for i := 2; i < len(jumps); i++ {
		dp[i] = int(math.Min(math.Abs(float64(jumps[i]-jumps[i-1]))+float64(dp[i-1]),
			math.Abs(float64(jumps[i]-jumps[i-2]))+float64(dp[i-2])))
		//fmt.Println("isi arr : ", dp[i-1], dp[i])
	}
	//fmt.Println(len(jumps) - 1)
	return dp[len(jumps)-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
