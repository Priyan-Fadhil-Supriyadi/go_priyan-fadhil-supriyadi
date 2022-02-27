package main

import (
	"fmt"
	"strconv"
)

func angkaMunculSekali(angka string) []int {
	berubah := []string{}
	for i := range angka {
		kata := string(angka[i])
		berubah = append(berubah, kata)
	}

	// sisa := []string{}
	pass := []string{}
	//visited := map[string]bool{}

	for i := 0; i < len(berubah); i++ {
		for j := i + 1; j < len(berubah); j++ {
			// fmt.Println("berubah : ", berubah[i], "apakah sama dengan = ", berubah[j])
			if berubah[i] == berubah[j] {
				berubah[j] = "NULL"
				break
			} else if j == len(berubah)-1 && berubah[i] != "NULL" {
				pass = append(pass, berubah[i])
			}
		}
	}

	var t2 = []int{}

	for _, i := range pass {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		t2 = append(t2, j)
	}
	return t2
}

func main() {
	fmt.Println(angkaMunculSekali("1234123"))
}
