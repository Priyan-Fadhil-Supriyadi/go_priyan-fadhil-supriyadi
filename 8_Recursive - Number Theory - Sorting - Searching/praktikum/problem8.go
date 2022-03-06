package main

import (
	"fmt"
)

type pair struct {
	name  string
	count int
}

func (v pair) String() string {
	return fmt.Sprintf("%s->%d", v.name, v.count)
}

func MostAppearItem(items []string) []pair {
	maps := make(map[string]int)

	sort := func(data []pair) []pair {
		for i := 0; i < len(data)-1; i++ {
			idx := i
			for j := i + 1; j < len(data); j++ {
				if data[idx].count > data[j].count {
					idx = j
				}
			}
			data[idx], data[i] = data[i], data[idx]
		}
		return data
	}

	for i := 0; i < len(items); i++ {
		maps[items[i]]++
	}

	var pr []pair
	var temp pair
	for key, value := range maps {
		temp.name = key
		temp.count = value
		pr = append(pr, temp)
	}
	pr = sort(pr)

	return pr
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}
