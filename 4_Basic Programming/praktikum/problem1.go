package main

import "fmt"

func main() {
	// input
	var T float64
	var r float64

	fmt.Print("tinggi tabung dan jari jari : ")
	fmt.Scanf("%f", &T)
	fmt.Scanf("%f", &r)

	// kode disini
	hasil := (2.0 * 3.14 * (r * r)) + (2.0 * 3.14 * r * T)
	fmt.Println("hasilnya adalah : ", hasil)
}
