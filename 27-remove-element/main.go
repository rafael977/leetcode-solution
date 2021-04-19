package main

import "fmt"

func removeElement(nums []int, val int) int {
	j := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}

	return j + 1
}

func main() {
	nums := []int{3, 2, 2, 3}
	fmt.Println(removeElement(nums, 2), nums)

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 2), nums)
}
