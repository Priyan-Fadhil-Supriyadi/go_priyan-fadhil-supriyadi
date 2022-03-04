package main

import (
	"fmt"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	// your code here
	var num int = 0
	a := make([]pair, 0, len(items))
	//b := pair{a[:]}
	//b := make([]pair, 0, len(items))
	//duplicate_frequency := make(map[string]int)

	visited := map[string]bool{}
	for i := 0; i < len(items); i++ {
		n := items[i]
		//fmt.Println(a)
		if visited[n] {
			if a[i].name == n {
				a = append(a, pair{items[i], num})
			}
			num += 1

			continue
		}

		a = append(a, pair{items[i], num})
		visited[n] = true
	}
	fmt.Println(a[1].name)
	return a
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	// golang->1 ruby->2 js->4
	// fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// // C->1 D->2 B->3 A->4
	// fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
	// // football->1 basketball->1 tenis->1
}
