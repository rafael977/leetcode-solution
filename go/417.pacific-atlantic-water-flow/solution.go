package pacificatlanticwaterflow

/*
 * @lc app=leetcode id=417 lang=golang
 *
 * [417] Pacific Atlantic Water Flow
 */

// @lc code=start
var dirs [][]int = [][]int{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func pacificAtlantic(heights [][]int) (ans [][]int) {
	m, n := len(heights), len(heights[0])
	po, ao := make([][]bool, m), make([][]bool, m)
	for i := range po {
		po[i] = make([]bool, n)
		ao[i] = make([]bool, n)
	}

	var dfs func(i, j int, ocean [][]bool)
	dfs = func(i, j int, ocean [][]bool) {
		ocean[i][j] = true
		for _, dir := range dirs {
			ii, jj := i+dir[0], j+dir[1]
			if ii >= 0 && ii < m && jj >= 0 && jj < n &&
				!ocean[ii][jj] &&
				heights[i][j] <= heights[ii][jj] {
				dfs(ii, jj, ocean)
			}
		}
	}

	for i := 0; i < m; i++ {
		dfs(i, 0, po)
		dfs(i, n-1, ao)
	}
	for j := 0; j < n; j++ {
		dfs(0, j, po)
		dfs(m-1, j, ao)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if po[i][j] && ao[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return
}

// @lc code=end
