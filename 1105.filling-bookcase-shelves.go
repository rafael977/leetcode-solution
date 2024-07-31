package leetcodesolution

/*
 * @lc app=leetcode id=1105 lang=golang
 *
 * [1105] Filling Bookcase Shelves
 */

// @lc code=start
func minHeightShelves(books [][]int, shelfWidth int) int {
	dp := make([]int, len(books)+1)
	dp[0] = 0
	for i := 0; i < len(books); i++ {
		t, h := books[i][0], books[i][1]
		dp[i+1] = h + dp[i]
		for j := i - 1; j >= 0; j-- {
			t += books[j][0]
			if t > shelfWidth {
				break
			}
			h = max(h, books[j][1])
			dp[i+1] = min(dp[i+1], dp[j]+h)
		}
	}

	return dp[len(books)]
}

// @lc code=end
