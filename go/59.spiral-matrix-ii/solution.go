package spiralmatrixii

/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 */

// @lc code=start
func generateMatrix(n int) (ans [][]int) {
	ans = make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	top, bottom, left, right := 0, n-1, 0, n-1
	i := 1
	for i <= n*n {
		for r, c := top, left; c <= right; c++ {
			ans[r][c] = i
			i++
		}
		top++

		for r, c := top, right; r <= bottom; r++ {
			ans[r][c] = i
			i++
		}
		right--

		for r, c := bottom, right; c >= left; c-- {
			ans[r][c] = i
			i++
		}
		bottom--

		for r, c := bottom, left; r >= top; r-- {
			ans[r][c] = i
			i++
		}
		left++
	}

	return
}

// @lc code=end
