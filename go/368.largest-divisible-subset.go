package leetcodesolution

import "sort"

/*
 * @lc app=leetcode id=368 lang=golang
 *
 * [368] Largest Divisible Subset
 */

// @lc code=start
func largestDivisibleSubset(nums []int) (ans []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	dp := make([]struct{ p, n int }, len(nums))
	for i := range dp {
		dp[i] = struct {
			p int
			n int
		}{p: i, n: 1}
	}
	for i, v := range nums {
		for j := 0; j < i; j++ {
			if nums[j]%v == 0 && dp[i].n <= 1+dp[j].n {
				dp[i].p = j
				dp[i].n = 1 + dp[j].n
			}
		}
	}
	max, i := 0, 0
	for j, v := range dp {
		if v.n > max {
			max = v.n
			i = j
		}
	}

	ans = append(ans, nums[i])
	for i != dp[i].p {
		i = dp[i].p
		ans = append(ans, nums[i])
	}
	return
}

// @lc code=end
