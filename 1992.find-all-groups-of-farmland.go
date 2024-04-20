package leetcodesolution

/*
 * @lc app=leetcode id=1992 lang=golang
 *
 * [1992] Find All Groups of Farmland
 */

// @lc code=start
func findFarmland(land [][]int) (ans [][]int) {
	m, n := len(land), len(land[0])
	for i, row := range land {
		for j, v := range row {
			if v == 0 || (i > 0 && land[i-1][j] != 0) ||
				(j > 0 && land[i][j-1] != 0) {
				continue
			}
			ii, jj := i, j
			for ii < m && land[ii][j] == 1 {
				ii++
			}
			for jj < n && land[i][jj] == 1 {
				jj++
			}
			ans = append(ans, []int{i, j, ii - 1, jj - 1})
		}
	}
	return
}

// @lc code=end
