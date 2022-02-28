package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	berubah := []string{}
	for i := range angka {
		kata := string(angka[i])
		berubah = append(berubah, kata)
	}

	pass := []string{}
	for i := 0; i < len(berubah); i++ {
		for j := i + 1; j < len(berubah); j++ {
			if berubah[i] == berubah[j] {
				berubah[j] = "NULL"
				break
			}
			if j == len(berubah)-1 && berubah[i] != "NULL" {
				pass = append(pass, berubah[i])
			}
			if i == len(berubah)-2 && berubah[j] != "NULL" {
				pass = append(pass, berubah[j])
			}
		}
	}

	var t = []int{}
	for _, i := range pass {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		t = append(t, j)
	}
	return t
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6, 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]

}
