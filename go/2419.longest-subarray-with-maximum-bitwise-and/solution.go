package longestsubarraywithmaximumbitwiseand

import "slices"

/*
 * @lc app=leetcode id=2419 lang=golang
 *
 * [2419] Longest Subarray With Maximum Bitwise AND
 */

// @lc code=start
func longestSubarray(nums []int) (ans int) {
	m := slices.Max(nums)
	s := -1
	for i, v := range nums {
		if v == m && s == -1 {
			s = i
			continue
		}
		if v != m && s != -1 {
			ans = max(ans, i-s)
			s = -1
			continue
		}
	}

	if s != -1 {
		ans = max(ans, len(nums)-s)
	}
	return
}

// @lc code=end
