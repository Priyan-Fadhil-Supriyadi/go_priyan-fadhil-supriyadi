package main

import "fmt"

func palindrome(input string) bool {
	// your code here
	kata := len(input) - 1
	sama := 0
	for i := range input {
		if input[i] == input[kata] {
			sama = sama + 1
		}
		kata = kata - 1
	}
	if sama == len(input) {
		return true
	}
	return false
}

func main() {
	fmt.Println(palindrome("civic"))       // true
	fmt.Println(palindrome("katak"))       // true
	fmt.Println(palindrome("kasur rusak")) // true
	fmt.Println(palindrome("mistar"))      // false
	fmt.Println(palindrome("lion"))        // false
}
