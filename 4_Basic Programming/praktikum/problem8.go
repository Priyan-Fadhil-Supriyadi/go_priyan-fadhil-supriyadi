package main

import "fmt"

func cetakTablePerkalian(number int) {
	for i := 1; i <= number; i++ {
		if number > 30 || number < 1 {
			fmt.Println("number merupakan bilangan 1 sampai 30")
			break
		}
		for j := 1; j <= number; j++ {
			num := j * i
			fmt.Print(num)
			if num <= 9 {
				fmt.Print(" ")
			} else if num > 99 {
				fmt.Print(" ")
				continue
			}
			fmt.Print("  ")
		}
		fmt.Println()
	}
}

func main() {
	cetakTablePerkalian(9)
}
