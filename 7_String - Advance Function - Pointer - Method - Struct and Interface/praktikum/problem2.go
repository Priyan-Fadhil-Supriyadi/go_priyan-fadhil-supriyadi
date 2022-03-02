package main

import "fmt"

func caesar(offset int, input string) string {
	// your code here
	ascii := []byte{}
	if offset > 26 {
		offset = offset % 26
	}
	for i := range input {
		ascii = append(ascii, input[i]+byte(offset))
		if ascii[i] > 122 {
			lebih := ascii[i] - 122
			ascii[i] = 96 + lebih
		}
	}

	var fix string
	for j := 0; j < len(ascii); j++ {
		fix = fix + string(ascii[j])
	}
	return fix

}

func main() {
	fmt.Println(caesar(3, "abc"))             // def
	fmt.Println(caesar(2, "alta"))            // cnvc
	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	// bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
	// mnopqrstuvwxyzabcdefghijkl
}
