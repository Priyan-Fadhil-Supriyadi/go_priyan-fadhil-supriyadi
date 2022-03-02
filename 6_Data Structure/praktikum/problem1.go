package main

import "fmt"

func ArrayMerge(arrayA, ArrayB []string) []string {
	arrayC := arrayA

	for i := 0; i < len(ArrayB); i++ {
		arrayC = append(arrayC, ArrayB[i])
	}
	sisa := []string{}
	visited := map[string]bool{}
	for i := 0; i < len(arrayC); i++ {
		n := arrayC[i]
		if visited[n] {
			continue
		}

		visited[n] = true
		sisa = append(sisa, arrayC[i])
	}
	return sisa
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))
}
