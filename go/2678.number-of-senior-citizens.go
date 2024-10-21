package leetcodesolution

/*
 * @lc app=leetcode id=2678 lang=golang
 *
 * [2678] Number of Senior Citizens
 */

// @lc code=start
func countSeniors(details []string) (ans int) {
	for _, d := range details {
		age := int(d[11]-'0')*10 + int(d[12]-'0')
		if age > 60 {
			ans++
		}
	}
	return
}

// @lc code=end
