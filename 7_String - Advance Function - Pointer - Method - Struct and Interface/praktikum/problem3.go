package main

import "fmt"

func swap(a, b *int) {
	var c int = *a
	*a = *b
	*b = c
}

func main() {
	a := 10
	b := 20

	swap(&a, &b)
	fmt.Println(a, b)
}
