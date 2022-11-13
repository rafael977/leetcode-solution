package reversevowelsofastring

/*
 * @lc app=leetcode id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 */

// @lc code=start
func reverseVowels(s string) string {
	buf := []byte(s)
	for i, j := 0, len(s)-1; i < j; {
		for i < len(s) && !isVowel(buf[i]) {
			i++
		}
		for j >= 0 && !isVowel(buf[j]) {
			j--
		}
		if i >= j {
			break
		}
		buf[i], buf[j] = buf[j], buf[i]
		i++
		j--
	}

	return string(buf)
}

func isVowel(c byte) bool {
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
		return true
	}
	return false
}

// @lc code=end
