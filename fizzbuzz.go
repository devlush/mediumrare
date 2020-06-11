
package main

import "fmt"

const N = 25

func build_terms() *[]string {

	cx := make([]string, N)

	cx[3]  = "Fizz"
	cx[5]  = "Buzz"
	cx[16] = "Jazz"

	return &cx
}

func build_factors() *map[int][]int {

	fx := make(map[int][]int)

	for i := 1; i <= N; i++ {
		//fx[i] = append(fx[i], 1)
		for d := 2; d*d <= i; d++ {
			if i%d == 0 {
				fx[i] = append(fx[i], d)
			}
		}
		for x := len(fx[i])-1; x >= 0; x-- {
			hf := i / fx[i][x]
			if hf * hf != i {
				fx[i] = append(fx[i], hf)
			}
		}
	}
	return &fx
}

func main() {
	cxref := build_terms()
	fxref := build_factors()
	//fmt.Println("\n", *cxref, "\n")
	//fmt.Println("\n\n", *fxref, "\n")

	for i := 1; i < N; i++ {
		fmt.Print(i, "\t")

		var complete = true
		for _, x := range (*fxref)[i] {
			if (*cxref)[x] == "" {
				complete = false
				break
			}
		}

		if complete {
			for _, x := range (*fxref)[i] {
				fmt.Print( (*cxref)[x] )
			}
		}

		if (*cxref)[i] != "" {
			fmt.Print( (*cxref)[i] )
		}
		fmt.Println()
	}
}

