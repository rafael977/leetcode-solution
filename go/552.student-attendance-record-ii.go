package leetcodesolution

/*
 * @lc app=leetcode id=552 lang=golang
 *
 * [552] Student Attendance Record II
 */

// @lc code=start
func checkRecord(n int) int {
	var mod int = (1e9 + 7)
	dp := make([][2][3]int, n+1)
	for i := range dp {
		for j := range dp[i] {
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	var dfs func(c, ac, lc int) int
	dfs = func(c, ac, lc int) int {
		if ac >= 2 || lc >= 3 {
			return 0
		}
		if c == 0 {
			return 1
		}
		if v := dp[c][ac][lc]; v != -1 {
			return v
		}
		r := (dfs(c-1, ac, 0)%mod + dfs(c-1, ac+1, 0)%mod + dfs(c-1, ac, lc+1)%mod) % mod
		dp[c][ac][lc] = r
		return r
	}

	return dfs(n, 0, 0)
}

// @lc code=end
