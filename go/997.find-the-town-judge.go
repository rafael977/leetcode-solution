package leetcodesolution

/*
 * @lc app=leetcode id=997 lang=golang
 *
 * [997] Find the Town Judge
 */

// @lc code=start
func findJudge(n int, trust [][]int) int {
	t := make([]int, n+1)
	for _, v := range trust {
		t[v[0]]--
		t[v[1]]++
	}
	for i := 1; i <= n; i++ {
		if t[i] == n-1 {
			return i
		}
	}
	return -1
}

// @lc code=end
