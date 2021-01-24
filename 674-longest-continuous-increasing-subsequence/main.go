package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	max := 0
	c := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			c++
		} else {
			c = 1
		}
		if max < c {
			max = c
		}
	}

	return max
}

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
	fmt.Println(findLengthOfLCIS([]int{2, 2, 2, 2}))
}
