package convert1darrayinto2darray

/*
 * @lc app=leetcode id=2022 lang=golang
 *
 * [2022] Convert 1D Array Into 2D Array
 */

// @lc code=start
func construct2DArray(original []int, m int, n int) (ans [][]int) {
	if len(original) != m*n {
		return
	}

	for i := 0; i < m; i++ {
		ans = append(ans, original[i*n:(i+1)*n])
	}
	return
}

// @lc code=end
