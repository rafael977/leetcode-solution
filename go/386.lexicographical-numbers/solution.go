package lexicographicalnumbers

/*
 * @lc app=leetcode id=386 lang=golang
 *
 * [386] Lexicographical Numbers
 */

// @lc code=start
func lexicalOrder(n int) (ans []int) {
	var dfs func(i int)
	dfs = func(i int) {
		if i > n {
			return
		}
		ans = append(ans, i)
		for j := 0; j <= 9; j++ {
			dfs(i*10 + j)
		}
	}
	for i := 1; i <= 9; i++ {
		dfs(i)
	}
	return
}

// @lc code=end
