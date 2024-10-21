package leetcodesolution

/*
 * @lc app=leetcode id=861 lang=golang
 *
 * [861] Score After Flipping Matrix
 */

// @lc code=start
func matrixScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	rows := make([]int, m)
	for i, r := range grid {
		rv := 0
		for _, v := range r {
			rv = rv<<1 + v
		}
		flip := (1<<n - 1) ^ rv
		if flip > rv {
			rows[i] = flip
		} else {
			rows[i] = rv
		}
	}

	for c := n - 1; c >= 0; c-- {
		sum := 0
		for _, r := range rows {
			sum += (r >> c) & 1
		}
		if sum <= m/2 {
			for i, r := range rows {
				rows[i] = r ^ (1 << c)
			}
		}
	}

	ans := 0
	for _, v := range rows {
		ans += v
	}
	return ans
}

// @lc code=end
