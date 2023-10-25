package integerbreak

/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 */

// @lc code=start
func integerBreak(n int) (ans int) {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	for k := 2; k <= n/2; k++ {
		p := 1
		c, r := n/k, n%k
		for i := 0; i < r; i++ {
			p *= (c + 1)
		}
		for i := 0; i < k-r; i++ {
			p *= c
		}
		ans = max(ans, p)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
