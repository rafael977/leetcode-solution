package maxareaofisland

/*
 * @lc app=leetcode id=695 lang=golang
 *
 * [695] Max Area of Island
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == -1 || j == -1 || i == m || j == n || grid[i][j] == 0 {
			return 0
		}

		grid[i][j] = 0
		return 1 + dfs(i-1, j) + dfs(i, j-1) + dfs(i+1, j) + dfs(i, j+1)
	}

	for i := range grid {
		for j := range grid[i] {
			ans = max(ans, dfs(i, j))
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
