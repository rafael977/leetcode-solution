package main

import (
	"fmt"
)

func numEquivDominoPairs(dominoes [][]int) int {
	counts := make([]int, 100)
	ans := 0
	for _, v := range dominoes {
		if v[0] > v[1] {
			v[0], v[1] = v[1], v[0]
		}
		i := v[0]*10 + v[1]
		ans += counts[i]
		counts[i]++
	}
	return ans
}

func main() {
	fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}))
}
