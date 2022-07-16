package outofboundarypaths

/*
 * @lc app=leetcode id=576 lang=golang
 *
 * [576] Out of Boundary Paths
 */

// @lc code=start
func findPaths(m int, n int, maxMove int, startRow int, startColumn int) (ans int) {
	const MOD = 1e9 + 7
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[startRow][startColumn] = 1

	for mv := 0; mv < maxMove; mv++ {
		dpNew := make([][]int, m)
		for i := range dpNew {
			dpNew[i] = make([]int, n)
		}

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if dp[i][j] == 0 {
					continue
				}
				for _, d := range dirs {
					ii, jj := i+d[0], j+d[1]
					if ii >= 0 && ii < m && jj >= 0 && jj < n {
						dpNew[ii][jj] = (dpNew[ii][jj] + dp[i][j]) % MOD
					} else {
						ans = (ans + dp[i][j]) % MOD
					}
				}
			}
		}

		dp = dpNew
	}

	return
}

// @lc code=end
