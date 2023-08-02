package countnegativenumbersinasortedmatrix

/*
 * @lc app=leetcode id=1351 lang=golang
 *
 * [1351] Count Negative Numbers in a Sorted Matrix
 */

// @lc code=start
func countNegatives(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	i, j := 0, n-1
	for i < m && j >= 0 {
		if grid[i][j] >= 0 {
			i++
		} else {
			ans += m - i
			j--
		}
	}
	return
}

// @lc code=end
