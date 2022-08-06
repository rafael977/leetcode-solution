package uniquepaths

/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */

// @lc code=start
func uniquePaths(m int, n int) (ans int) {
	ans = 1
	for x, y := n, 1; y <= m-1; x, y = x+1, y+1 {
		ans = ans * x / y
	}
	return
}

// @lc code=end
