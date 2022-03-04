package main

import (
	"fmt"
)

func playingDomino(cards [][]int, deck []int) interface{} {
	// your code here
	var max int
	var res []int
	contains := func(data []int, check int) bool {
		for _, val := range data {
			if val == check {
				return true
			}
		}
		return false
	}

	for _, val := range cards {
		for _, value := range val {
			if contains(deck, value) {
				if max < val[0]+val[1] {
					max = val[0] + val[1]
					res = append(res, val[0])
					res = append(res, val[1])
				}
				break
			}
		}
	}

	if max != 0 {
		return res
	}
	return "tutup kartu"
}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))
	// [3, 4]
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))
	// [6 5]
	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))
	// "tutup kartu"
}
