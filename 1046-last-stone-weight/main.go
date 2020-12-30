package main

import (
	"fmt"
	"sort"
)

func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		l := len(stones)
		sort.Ints(stones)
		x, y := stones[l-2], stones[l-1]
		stones = append(stones[:l-2], y-x)
	}

	return stones[0]
}

func main() {
	arr := []int{2, 7, 4, 1, 8, 1}
	fmt.Printf("%d", lastStoneWeight(arr))
}
