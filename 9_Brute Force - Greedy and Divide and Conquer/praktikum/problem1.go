package main

import "fmt"

func SimpleEquations(a, b, c int) {
	// your code here
	var x, y, z int
	for x = 1; a-(x+y+z) >= 0 || b-(x*y*z) >= 0 || c-(x*x+y*y+z*z) >= 0; x++ {
		for y = 1; a-(x+y+z) >= 0 || b-(x*y*z) >= 0 || c-(x*x+y*y+z*z) >= 0; y++ {
			for z = 1; a-(x+y+z) >= 0 || b-(x*y*z) >= 0 || c-(x*x+y*y+z*z) >= 0; z++ {
				if a-(x+y+z) == 0 && b-(x*y*z) == 0 && c-(x*x+y*y+z*z) == 0 {
					fmt.Println(x, y, z)
					return
				}
			}
			z = 1
		}
		y = 1
	}
	fmt.Println("no solution")
}

func main() {
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
}
