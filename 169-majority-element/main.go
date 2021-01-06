package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Printf("%v\n", majorityElement(nums))

	nums = []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Printf("%v\n", majorityElement(nums))
}
