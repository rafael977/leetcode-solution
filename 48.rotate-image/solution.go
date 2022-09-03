package rotateimage

/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */

// @lc code=start
func rotate(matrix [][]int) {
	n := len(matrix)
	arr := make([]int, n*n)
	for i, row := range matrix {
		for j, v := range row {
			arr[i*n+j] = v
		}
	}

	i, j := 0, n-1
	for _, v := range arr {
		matrix[i][j] = v
		if i+1 >= n {
			j--
			i = 0
		} else {
			i++
		}
	}
}

// @lc code=end
