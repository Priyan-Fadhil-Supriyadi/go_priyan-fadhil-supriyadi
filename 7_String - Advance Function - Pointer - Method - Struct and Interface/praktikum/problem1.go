package main

import (
	"fmt"
)

func Compare(a, b string) string {
	// your code here
	var panjang, kata, fix string
	if len(a) < len(b) {
		kata = a
		panjang = b
	} else {
		kata = b
		panjang = a
	}

	for i := 0; i < len(panjang); i++ {
		substring := panjang[0:i]
		if substring == kata {
			fix = substring
		}
	}
	return fix
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}
