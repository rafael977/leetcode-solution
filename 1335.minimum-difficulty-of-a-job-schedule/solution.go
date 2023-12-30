package minimumdifficultyofajobschedule

import "math"

/*
 * @lc app=leetcode id=1335 lang=golang
 *
 * [1335] Minimum Difficulty of a Job Schedule
 */

// @lc code=start
func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}
	rangeMax := make([][]int, n)
	for i := range rangeMax {
		rangeMax[i] = make([]int, n)
		rangeMax[i][i] = jobDifficulty[i]
	}
	for l := 0; l < n; l++ {
		for r := l + 1; r < n; r++ {
			rangeMax[l][r] = max(rangeMax[l][r-1], jobDifficulty[r])
		}
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, d+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}
	for i := range dp {
		dp[i][1] = rangeMax[0][i]
	}
	for i := 0; i < n; i++ {
		for k := 2; k <= min(i+1, d); k++ {
			for j := k - 2; j < i; j++ {
				dp[i][k] = min(dp[i][k], dp[j][k-1]+rangeMax[j+1][i])
			}
		}
	}
	return dp[n-1][d]
}

// @lc code=end
