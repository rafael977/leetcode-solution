package leetcodesolution

/*
 * @lc app=leetcode id=2864 lang=golang
 *
 * [2864] Maximum Odd Binary Number
 */

// @lc code=start
func maximumOddBinaryNumber(s string) string {
	sb := []byte(s)
	i, j := 0, 0
	for i < len(sb) {
		for j < len(sb) && sb[j] == '0' {
			j++
		}
		if j == len(sb) {
			break
		}
		sb[i], sb[j] = sb[j], sb[i]
		i++
		j++
	}

	sb[i-1], sb[len(sb)-1] = sb[len(sb)-1], sb[i-1]
	return string(sb)
}

// @lc code=end
