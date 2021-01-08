package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n

	l, r := make([]int, n-k), make([]int, k)
	copy(l, nums[:n-k])
	copy(r, nums[n-k:])

	copy(nums, append(r, l...))
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Printf("%v\n", nums)

	nums = []int{1}
	rotate(nums, 2)
	fmt.Printf("%v\n", nums)
}
