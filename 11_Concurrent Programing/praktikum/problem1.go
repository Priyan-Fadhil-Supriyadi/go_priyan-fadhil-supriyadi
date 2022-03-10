package main

import (
	"fmt"
	"time"
)

type mapping map[string]int

func Frekuensi(s string) mapping {
	m := mapping{}
	for _, r := range s {
		m[string(r)]++
	}
	return m
}

func Hitung(k string) {
	channel := make(chan mapping, 2)

	wordSplit := []string{k[:len(k)/2], k[len(k)/2:]}
	for _, val := range wordSplit {
		go func(w string) {
			channel <- Frekuensi(w)
			time.Sleep(1000 * time.Millisecond)
		}(val)
	}
	hasil := mapping{}
	for range wordSplit {
		for key, value := range <-channel {
			hasil[key] += value
		}
	}
	for i, v := range hasil {
		fmt.Println(i, ":", v)
	}
}

func main() {
	// kode disini
	kalimat := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
	Hitung(kalimat)
}
