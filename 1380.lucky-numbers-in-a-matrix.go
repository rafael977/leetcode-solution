package leetcodesolution

/*
 * @lc app=leetcode id=1380 lang=golang
 *
 * [1380] Lucky Numbers in a Matrix
 */

// @lc code=start
func luckyNumbers(matrix [][]int) (ans []int) {
	m, n := len(matrix), len(matrix[0])
	rowMin, colMax := make([]int, m), make([]int, n)
	for i, r := range matrix {
		for j, v := range r {
			if rowMin[i] == 0 {
				rowMin[i] = v
			}
			rowMin[i] = min(rowMin[i], v)
			colMax[j] = max(colMax[j], v)
		}
	}
	for i, r := range matrix {
		for j, v := range r {
			if v == rowMin[i] && v == colMax[j] {
				ans = append(ans, v)
			}
		}
	}
	return
}

// @lc code=end
