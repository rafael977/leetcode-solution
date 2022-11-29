package wordsearch

/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if k == len(word) {
			return true
		}
		if i < 0 || i >= m || j < 0 || j >= n ||
			vis[i][j] || board[i][j] != word[k] {
			return false
		}
		vis[i][j] = true
		defer func() {
			vis[i][j] = false
		}()
		return dfs(i-1, j, k+1) || dfs(i+1, j, k+1) ||
			dfs(i, j-1, k+1) || dfs(i, j+1, k+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// @lc code=end
