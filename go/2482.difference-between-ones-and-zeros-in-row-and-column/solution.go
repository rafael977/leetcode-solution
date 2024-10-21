package differencebetweenonesandzerosinrowandcolumn

/*
 * @lc app=leetcode id=2482 lang=golang
 *
 * [2482] Difference Between Ones and Zeros in Row and Column
 */

// @lc code=start
func onesMinusZeros(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	row, col := make([]int, m), make([]int, n)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				row[i]--
				col[j]--
			} else {
				row[i]++
				col[j]++
			}
		}
	}

	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans[i][j] = row[i] + col[j]
		}
	}

	return ans
}

// @lc code=end
