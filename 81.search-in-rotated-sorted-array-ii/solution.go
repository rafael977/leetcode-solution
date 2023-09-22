package searchinrotatedsortedarrayii

import (
	"sort"
)

/*
 * @lc app=leetcode id=81 lang=golang
 *
 * [81] Search in Rotated Sorted Array II
 */

// @lc code=start
func search(nums []int, target int) bool {
	n := len(nums)
	k := 0
	for i := range nums {
		if i == 0 {
			continue
		}
		if nums[i] < nums[i-1] {
			k = i
		}
	}
	x := sort.Search(n, func(i int) bool {
		p := (i + k) % n
		return nums[p] >= target
	})

	return x < n && nums[(x+k)%n] == target
}

// @lc code=end
