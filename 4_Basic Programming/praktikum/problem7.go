package main

import "fmt"

func playWithAsterik(n int) {
	num := n - 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= num; j++ {
			fmt.Printf(" ")
		}
		num = num - 1
		for k := 1; k <= i; k++ {
			fmt.Printf("*")
			fmt.Printf(" ")
		}
		fmt.Println()
	}
}

func main() {
	playWithAsterik(5)
}
