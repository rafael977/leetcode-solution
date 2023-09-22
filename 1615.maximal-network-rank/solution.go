package maximalnetworkrank

/*
 * @lc app=leetcode id=1615 lang=golang
 *
 * [1615] Maximal Network Rank
 */

// @lc code=start
func maximalNetworkRank(n int, roads [][]int) int {
	deg := make([]int, n)
	conn := make([][]int, n)
	for i := range conn {
		conn[i] = make([]int, n)
	}

	for _, r := range roads {
		i, j := r[0], r[1]
		deg[i]++
		deg[j]++
		conn[i][j] = 1
		conn[j][i] = 1
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ans = max(ans, deg[i]+deg[j]-conn[i][j])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
