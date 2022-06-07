package transposematrix

/*
 * @lc app=leetcode id=867 lang=golang
 *
 * [867] Transpose Matrix
 */

// @lc code=start
func transpose(matrix [][]int) (tm [][]int) {
	m, n := len(matrix), len(matrix[0])
	tm = make([][]int, n)
	for i := range tm {
		tm[i] = make([]int, m)
		for j := range tm[i] {
			tm[i][j] = matrix[j][i]
		}
	}

	return
}

// @lc code=end
