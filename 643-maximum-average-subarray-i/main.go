package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	max := sum
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		if max < sum {
			max = sum
		}
	}

	return float64(max) / float64(k)
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}
