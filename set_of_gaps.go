/*
Given the coefficients of a Diophantine equation,
find the set of elements for which it has no solution
*/

package main

import "fmt"

func contains(array []int, elemento int) bool {
	for _, valor := range array {
		if valor == elemento {
			return true
		}
	}
	return false
}

func set_of_gaps(a int, b int) []int {
	var set []int

	for m := 1; m < b; m++ {
		var res int = 1
		var n int = 1

		for res > 0 {
			res = a*m - 7*n

			if res > 0 {
				set = append(set, res)
			}

			n += 1
		}
	}

	return set
}

func main() {
	fmt.Println(contains(set_of_gaps(11, 7), 100))
}
