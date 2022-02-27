package main

import "fmt"

func PairSum(arr []int, target int) []int {
	indeks := []int{}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			// fmt.Println("berubah : ", berubah[i], "apakah sama dengan = ", berubah[j])
			sum := arr[i] + arr[j]
			if sum == target {
				indeks = append(indeks, i)
				indeks = append(indeks, j)
			}
		}
	}
	return indeks
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))
	/*
	   [1, 3]
	*/
}
