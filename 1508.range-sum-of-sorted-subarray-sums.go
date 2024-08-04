package leetcodesolution

import "sort"

/*
 * @lc app=leetcode id=1508 lang=golang
 *
 * [1508] Range Sum of Sorted Subarray Sums
 */

// @lc code=start
func rangeSum(nums []int, n int, left int, right int) (ans int) {
	var MOD int = 1e9 + 7
	sums := make([]int, 0)
	for i := 0; i < n; i++ {
		s := nums[i]
		sums = append(sums, s)
		for j := i + 1; j < n; j++ {
			s += nums[j]
			sums = append(sums, s)
		}
	}

	sort.Ints(sums)
	for i := left - 1; i < right; i++ {
		ans = (ans + sums[i]) % MOD
	}
	return
}

// @lc code=end
