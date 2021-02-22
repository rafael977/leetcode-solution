package main

import "fmt"

func checkPossibility(nums []int) bool {
	c := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			c--
			if c < 0 || (i-2 >= 0 && nums[i] < nums[i-2]) {
				return false
			}
		}
	}

	return true
}

func main() {
	// fmt.Println(checkPossibility([]int{4, 2, 3}))
	// fmt.Println(checkPossibility([]int{4, 3, 2}))
	// fmt.Println(checkPossibility([]int{3, 4, 2, 3}))
	fmt.Println(checkPossibility([]int{5, 7, 1, 8}))
}
