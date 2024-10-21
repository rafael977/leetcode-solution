package decodedstringatindex

/*
 * @lc app=leetcode id=880 lang=golang
 *
 * [880] Decoded String at Index
 */

// @lc code=start
func decodeAtIndex(s string, k int) string {
	n := 0
	for i := range s {
		if isLetter(s[i]) {
			n++
		}
		if isDigit(s[i]) {
			n *= int(s[i] - '0')
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		k %= n
		if k == 0 && isLetter(s[i]) {
			return string(s[i])
		}

		if isDigit(s[i]) {
			n /= int(s[i] - '0')
		}
		if isLetter(s[i]) {
			n--
		}
	}
	return ""
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isLetter(c byte) bool {
	return c >= 'a' && c <= 'z'
}

// @lc code=end
