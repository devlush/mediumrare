
package main

import "fmt"

const N = 25	// max limit for factors

func build_terms() *[]string {

	cx := make([]string, N)

	cx[3]  = "Fizz"
	cx[5]  = "Buzz"
	cx[16] = "Jazz"

	return &cx
}

func build_factors() *map[int][]int {

	// map of integer arrays
	fx := make(map[int][]int)

	for i := 1; i <= N; i++ {
	// loop over integers 1 to N
	// calculate the possible factors for each integer
	// and store as an integer array
		for d := 2; d*d <= i; d++ {
		// loop over integer divisors from
		// zero up to the square root of i (product)
			if i%d == 0 {
				fx[i] = append(fx[i], d)
			}
		}
		// for every factor found below the sqrt,
		// divide it from the product (i) and store
		// the quotient as an additional factor
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

		for _, x := range (*fxref)[i] {
			fmt.Print( (*cxref)[x] )
		}

		if (*cxref)[i] != "" {
			fmt.Print( (*cxref)[i] )
		}
		fmt.Println()
	}
}

