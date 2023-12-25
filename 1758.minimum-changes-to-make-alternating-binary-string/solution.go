package minimumchangestomakealternatingbinarystring

/*
 * @lc app=leetcode id=1758 lang=golang
 *
 * [1758] Minimum Changes To Make Alternating Binary String
 */

// @lc code=start
func minOperations(s string) int {
	isOne := false
	c0 := 0
	for _, v := range s {
		if isOne && v == '0' || (!isOne && v == '1') {
			c0++
		}
	}

	return min(change(s, '0'), change(s, '1'))
}

func flip(c rune) rune {
	if c == '0' {
		return '1'
	}
	return '0'
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func change(s string, b rune) (c int) {
	for _, v := range s {
		if v != b {
			c++
		}
		b = flip(b)
	}
	return
}

// @lc code=end
