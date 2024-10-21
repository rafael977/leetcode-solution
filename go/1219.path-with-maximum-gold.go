package leetcodesolution

/*
 * @lc app=leetcode id=1219 lang=golang
 *
 * [1219] Path with Maximum Gold
 */

// @lc code=start
func getMaximumGold(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	var dfs func(i, j, sum int)
	dfs = func(i, j, sum int) {
		v := grid[i][j]
		sum += v
		ans = max(ans, sum)

		grid[i][j] = 0
		for _, d := range dirs {
			ii, jj := i+d[0], j+d[1]
			if ii >= 0 && ii < m && jj >= 0 && jj < n && grid[ii][jj] != 0 {
				dfs(ii, jj, sum)
			}
		}
		grid[i][j] = v
	}

	for i := range grid {
		for j, v := range grid[i] {
			if v != 0 {
				dfs(i, j, 0)
			}
		}
	}
	return
}

// @lc code=end
