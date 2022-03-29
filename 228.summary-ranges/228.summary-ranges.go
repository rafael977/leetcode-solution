package main

import "fmt"

/*
 * @lc app=leetcode id=228 lang=golang
 *
 * [228] Summary Ranges
 */

// @lc code=start
func summaryRanges(nums []int) (ans []string) {
	s := 0
	for c := 1; c <= len(nums); c++ {
		if c == len(nums) || nums[c-1]+1 != nums[c] {
			if s == c-1 {
				ans = append(ans, fmt.Sprintf("%d", nums[s]))
			} else {
				ans = append(ans, fmt.Sprintf("%d->%d", nums[s], nums[c-1]))
			}
			s = c
		}
	}
	return
}

// @lc code=end
