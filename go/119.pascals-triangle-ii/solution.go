package pascalstriangleii

/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */

// @lc code=start
func getRow(rowIndex int) (ans []int) {
	ans = append(ans, 1)
	n := 1
	for k := 1; k <= rowIndex; k++ {
		n = n * (rowIndex - k + 1) / k
		ans = append(ans, n)
	}
	return
}

// @lc code=end
