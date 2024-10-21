package main

import "fmt"

func pivotIndex(nums []int) int {
	l, r := 0, 0
	for i := 1; i < len(nums); i++ {
		r += nums[i]
	}

	if l == r {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		l += nums[i-1]
		r -= nums[i]

		if l == r {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
}
