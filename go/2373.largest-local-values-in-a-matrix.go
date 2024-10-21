package leetcodesolution

/*
 * @lc app=leetcode id=2373 lang=golang
 *
 * [2373] Largest Local Values in a Matrix
 */

// @lc code=start
func largestLocal(grid [][]int) [][]int {
	maskMax := func(x, y int) int {
		m := 0
		for i := x - 1; i <= x+1; i++ {
			for j := y - 1; j <= y+1; j++ {
				m = max(m, grid[i][j])
			}
		}
		return m
	}

	n := len(grid)
	ans := make([][]int, n-2)
	for i := range ans {
		ans[i] = make([]int, n-2)
	}
	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			ans[i-1][j-1] = maskMax(i, j)
		}
	}
	return ans
}

// @lc code=end
