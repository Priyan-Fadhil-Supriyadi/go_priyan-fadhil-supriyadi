package main

import (
	"fmt"
)

func primeNumber(number int) bool {
	// your code here
	var prima bool = true
	for i := 2; i < number; i++ {
		if number%i == 0 {
			prima = false
		}
	}
	// if prima == true {
	// 	fmt.Println("Bilangan Prima")
	// } else {
	// 	fmt.Println("Bukan Bilangan Prima")
	// }
	return prima
}

func main() {
	fmt.Println(primeNumber(11))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}
