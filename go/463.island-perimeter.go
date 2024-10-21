package leetcodesolution

/*
 * @lc app=leetcode id=463 lang=golang
 *
 * [463] Island Perimeter
 */

// @lc code=start
var dirs = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func islandPerimeter(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	i, j := 0, 0
	for i < m {
		found := false
		for j < n {
			if grid[i][j] == 1 {
				found = true
				break
			}
			j++
		}
		if found {
			break
		}
		j = 0
		i++
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 2 {
			return
		}
		grid[i][j] = 2
		for _, d := range dirs {
			ii, jj := i+d[0], j+d[1]
			if ii < 0 || ii >= m || jj < 0 || jj >= n || grid[ii][jj] == 0 {
				ans++
				continue
			} else if grid[ii][jj] == 1 {
				dfs(ii, jj)
			}
		}
	}

	dfs(i, j)
	return
}

// @lc code=end
