package longestarithmeticsubsequenceofgivendifference

/*
 * @lc app=leetcode id=1218 lang=golang
 *
 * [1218] Longest Arithmetic Subsequence of Given Difference
 */

// @lc code=start
func longestSubsequence(arr []int, difference int) (ans int) {
	dp := map[int]int{}
	for _, v := range arr {
		dp[v] = dp[v-difference] + 1
		if dp[v] > ans {
			ans = dp[v]
		}
	}
	return
}

// @lc code=end
