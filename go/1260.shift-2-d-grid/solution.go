package shift2dgrid

/*
 * @lc app=leetcode id=1260 lang=golang
 *
 * [1260] Shift 2D Grid
 */

// @lc code=start
func shiftGrid(grid [][]int, k int) (ans [][]int) {
	m, n := len(grid), len(grid[0])
	c := m * n

	ans = make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			x := (i*n + j - k) % c
			for x < 0 {
				x += c
			}
			ii, jj := x/n, x%n
			ans[i][j] = grid[ii][jj]
		}
	}

	return
}

// @lc code=end
