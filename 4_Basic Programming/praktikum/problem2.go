package main

import "fmt"

func main() {
	// input
	var studentScore int
	var siswa string
	fmt.Print("Nama Siswa : ")
	fmt.Scanln(&siswa)
	fmt.Print("Nilai Siswa : ")
	fmt.Scanln(&studentScore)

	// Precess: Your Solution Code Here
	fmt.Print(siswa, " mendapatkan ")
	if studentScore < 0 && studentScore > 100 {
		fmt.Println("Nilai Invalid")
	} else if studentScore >= 80 {
		fmt.Println("Nilai A")
	} else if studentScore >= 65 && studentScore <= 79 {
		fmt.Println("Nilai B")
	} else if studentScore >= 50 && studentScore <= 64 {
		fmt.Println("Nilai C")
	} else if studentScore >= 35 && studentScore <= 49 {
		fmt.Println("Nilai D")
	} else if studentScore <= 34 {
		fmt.Println("Nilai E")
	}

	// Output
	// Nilai A
}
