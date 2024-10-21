package leetcodesolution

import "sort"

/*
 * @lc app=leetcode id=2191 lang=golang
 *
 * [2191] Sort the Jumbled Numbers
 */

// @lc code=start
func sortJumbled(mapping []int, nums []int) []int {
	mn := make(map[int]int)
	for _, v := range nums {
		mul := 1
		x, r := 0, v
		if v == 0 {
			mn[v] = mapping[0]
			continue
		}
		for r > 0 {
			x += mul * mapping[r%10]
			r = r / 10
			mul *= 10
		}
		mn[v] = x
	}

	sort.SliceStable(nums, func(i, j int) bool {
		return mn[nums[i]] < mn[nums[j]]
	})
	return nums
}

// @lc code=end
