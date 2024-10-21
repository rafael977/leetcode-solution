package minimumbitflipstoconvertnumber

/*
 * @lc app=leetcode id=2220 lang=golang
 *
 * [2220] Minimum Bit Flips to Convert Number
 */

// @lc code=start
func minBitFlips(start int, goal int) (ans int) {
	r := start ^ goal
	for r > 0 {
		ans += r & 1
		r = r >> 1
	}
	return
}

// @lc code=end
