package specialpositionsinabinarymatrix

/*
 * @lc app=leetcode id=1582 lang=golang
 *
 * [1582] Special Positions in a Binary Matrix
 */

// @lc code=start
func numSpecial(mat [][]int) (ans int) {
	m, n := len(mat), len(mat[0])
	row, col := make([]int, m), make([]int, n)
	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 1 {
				row[i]++
				col[j]++
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 && row[i] == 1 && col[j] == 1 {
				ans++
			}
		}
	}
	return
}

// @lc code=end
