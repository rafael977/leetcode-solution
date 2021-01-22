package main

import "fmt"

func addToArrayForm(A []int, K int) []int {
	r := make([]int, len(A))
	for i := len(A) - 1; i >= 0; i-- {
		s := A[i] + K
		r[i] = s % 10
		K = s / 10
	}

	for K > 0 {
		r = append([]int{K % 10}, r...)
		K = K / 10
	}

	return r
}

func main() {
	fmt.Printf("%v\n", addToArrayForm([]int{1, 2, 0, 0}, 34))

	fmt.Printf("%v\n", addToArrayForm([]int{2, 7, 4}, 181))

	fmt.Printf("%v\n", addToArrayForm([]int{2, 1, 5}, 806))

	fmt.Printf("%v\n", addToArrayForm([]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 1))

	fmt.Printf("%v\n", addToArrayForm([]int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 0, 6, 3}, 516))

	fmt.Printf("%v\n", addToArrayForm([]int{0}, 1000))
}
