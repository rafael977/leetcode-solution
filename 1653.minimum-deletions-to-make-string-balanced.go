package leetcodesolution

/*
 * @lc app=leetcode id=1653 lang=golang
 *
 * [1653] Minimum Deletions to Make String Balanced
 */

// @lc code=start
func minimumDeletions(s string) (ans int) {
	righta := 0
	for _, c := range s {
		if c == 'a' {
			righta++
		}
	}
	ans = righta
	leftb := 0
	for _, c := range s {
		if c == 'a' {
			righta--
		} else {
			leftb++
		}
		ans = min(ans, leftb+righta)
	}
	return
}

// @lc code=end
