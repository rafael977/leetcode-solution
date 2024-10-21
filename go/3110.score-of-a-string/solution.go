package leetcodesolution

/*
 * @lc app=leetcode id=3110 lang=golang
 *
 * [3110] Score of a String
 */

// @lc code=start
func scoreOfString(s string) (ans int) {
	for i := 1; i < len(s); i++ {
		d := int(s[i]) - int(s[i-1])
		if d < 0 {
			d = -d
		}
		ans += d
	}
	return
}

// @lc code=end
