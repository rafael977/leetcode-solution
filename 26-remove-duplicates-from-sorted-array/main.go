package main

import (
	"fmt"
)

// func removeDuplicates(nums []int) int {
// 	ans := 0
// 	cur := math.MinInt32

// 	for _, v := range nums {
// 		if cur != v {
// 			nums[ans] = v
// 			cur = v
// 			ans++
// 		}
// 	}

// 	return ans
// }

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	ans := 0

	for _, v := range nums {
		if nums[ans] != v {
			ans++
			nums[ans] = v
		}
	}

	return ans + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	n := removeDuplicates(nums)
	nums = nums[:n]
	fmt.Println(n, nums)

	nums = []int{}
	n = removeDuplicates(nums)
	fmt.Println(n, nums)
}
