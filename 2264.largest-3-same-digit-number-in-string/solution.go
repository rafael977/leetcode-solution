package largest3samedigitnumberinstring

import "fmt"

/*
 * @lc app=leetcode id=2264 lang=golang
 *
 * [2264] Largest 3-Same-Digit Number in String
 */

// @lc code=start
func largestGoodInteger(num string) string {
	ans := ""
	for i := 0; i < len(num)-2; i++ {
		c := num[i]
		if num[i+1] == c && num[i+2] == c && (ans == "" || ans < fmt.Sprintf("%c%c%c", c, c, c)) {
			ans = fmt.Sprintf("%c%c%c", c, c, c)
		}
	}
	return ans
}

// @lc code=end
