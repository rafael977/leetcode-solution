package spiralmatrix

/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
var dirs [][]int = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func spiralOrder(matrix [][]int) (ans []int) {
	m, n := len(matrix), len(matrix[0])
	u, b, l, r := 0, m-1, 0, n-1
	for {
		for i := l; i <= r; i++ {
			ans = append(ans, matrix[u][i])
		}
		u++
		if u > b {
			break
		}
		for j := u; j <= b; j++ {
			ans = append(ans, matrix[j][r])
		}
		r--
		if l > r {
			break
		}
		for i := r; i >= l; i-- {
			ans = append(ans, matrix[b][i])
		}
		b--
		if u > b {
			break
		}
		for j := b; j >= u; j-- {
			ans = append(ans, matrix[j][l])
		}
		l++
		if l > r {
			break
		}
	}
	return
}

// @lc code=end
