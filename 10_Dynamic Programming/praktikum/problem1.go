package main

import "fmt"

var memory map[int]int = map[int]int{}

func fibo(n int) int {
	// your code here
	if n <= 1 {
		memory[n] = n
	} else if _, exist := memory[n]; !exist {
		memory[n] = fibo(n-1) + fibo(n-2)
	}
	return memory[n]
}

func main() {
	fmt.Println(fibo(0))  // 0
	fmt.Println(fibo(1))  // 1
	fmt.Println(fibo(2))  // 1
	fmt.Println(fibo(3))  // 2
	fmt.Println(fibo(5))  // 5
	fmt.Println(fibo(6))  // 8
	fmt.Println(fibo(7))  // 13
	fmt.Println(fibo(9))  // 13
	fmt.Println(fibo(10)) // 55
}
