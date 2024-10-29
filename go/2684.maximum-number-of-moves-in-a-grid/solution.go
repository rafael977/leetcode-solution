// Created by Rafael Shen at 2024/10/29 21:12
// leetgo: 1.4.9
// https://leetcode.com/problems/maximum-number-of-moves-in-a-grid/

package maximumnumberofmovesinagrid

// @lc code=begin

func maxMoves(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	q := make(map[int]bool)
	for i := 0; i < m; i++ {
		q[i] = true
	}

	for j := 1; j < n; j++ {
		nq := make(map[int]bool)
		for i := range q {
			for ii := i - 1; ii <= i+1; ii++ {
				if ii >= 0 && ii < m && grid[i][j-1] < grid[ii][j] {
					nq[ii] = true
				}
			}
		}
		q = nq
		if len(q) == 0 {
			return j - 1
		}
	}
	return n - 1
}

// @lc code=end
