package deletecolumnstomakesorted

/*
 * @lc app=leetcode id=944 lang=golang
 *
 * [944] Delete Columns to Make Sorted
 */

// @lc code=start
func minDeletionSize(strs []string) (ans int) {
	n := len(strs)
	for i := range strs[0] {
		for j := 1; j < n; j++ {
			if strs[j-1][i] > strs[j][i] {
				ans++
				break
			}
		}
	}
	return
}

// @lc code=end
