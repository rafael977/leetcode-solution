package numberofislands

/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslands(grid [][]byte) (ans int) {
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		grid[i][j] = '0'

		if i-1 >= 0 && grid[i-1][j] == '1' {
			dfs(i-1, j)
		}
		if i+1 < m && grid[i+1][j] == '1' {
			dfs(i+1, j)
		}
		if j-1 >= 0 && grid[i][j-1] == '1' {
			dfs(i, j-1)
		}
		if j+1 < n && grid[i][j+1] == '1' {
			dfs(i, j+1)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return
}

// @lc code=end
