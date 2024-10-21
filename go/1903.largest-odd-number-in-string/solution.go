package largestoddnumberinstring

/*
 * @lc app=leetcode id=1903 lang=golang
 *
 * [1903] Largest Odd Number in String
 */

// @lc code=start
func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		if num[i] == '1' ||
			num[i] == '3' ||
			num[i] == '5' ||
			num[i] == '7' ||
			num[i] == '9' {
			return num[:i+1]
		}
	}
	return ""
}

// @lc code=end
