package minimumfallingpathsum

/*
 * @lc app=leetcode id=931 lang=golang
 *
 * [931] Minimum Falling Path Sum
 */

// @lc code=start
func minFallingPathSum(matrix [][]int) (ans int) {
	n := len(matrix)

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < n; j++ {
			best := matrix[i+1][j]
			if j-1 >= 0 {
				best = min(best, matrix[i+1][j-1])
			}
			if j+1 < n {
				best = min(best, matrix[i+1][j+1])
			}
			matrix[i][j] += best
		}
	}

	ans = 10001
	for j := 0; j < n; j++ {
		ans = min(ans, matrix[0][j])
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
