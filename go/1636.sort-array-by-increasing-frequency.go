package leetcodesolution

import "sort"

/*
 * @lc app=leetcode id=1636 lang=golang
 *
 * [1636] Sort Array by Increasing Frequency
 */

// @lc code=start
func frequencySort(nums []int) []int {
	f := make(map[int]int)
	for _, v := range nums {
		f[v]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if f[nums[i]] == f[nums[j]] {
			return nums[i] > nums[j]
		}
		return f[nums[i]] < f[nums[j]]
	})
	return nums
}

// @lc code=end
