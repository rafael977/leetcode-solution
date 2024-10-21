package separateblackandwhiteballs

/*
 * @lc app=leetcode id=2938 lang=golang
 *
 * [2938] Separate Black and White Balls
 */

// @lc code=start
func minimumSteps(s string) int64 {
	ans := 0
	l, r := 0, len(s)-1
	for l < r {
		for l < len(s) && s[l] == '0' {
			l++
		}
		for r >= 0 && s[r] == '1' {
			r--
		}
		if l >= r {
			break
		}
		ans += r - l
		l++
		r--
	}
	return int64(ans)
}

// @lc code=end
