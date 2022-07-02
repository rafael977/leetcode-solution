package partitioningintominimumnumberofdecibinarynumbers

/*
 * @lc app=leetcode id=1689 lang=golang
 *
 * [1689] Partitioning Into Minimum Number Of Deci-Binary Numbers
 */

// @lc code=start
func minPartitions(n string) (ans int) {
	for i := range n {
		ans = max(ans, int(n[i]-'0'))
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
