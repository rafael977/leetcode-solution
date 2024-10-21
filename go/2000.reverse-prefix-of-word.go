package leetcodesolution

/*
 * @lc app=leetcode id=2000 lang=golang
 *
 * [2000] Reverse Prefix of Word
 */

// @lc code=start
func reversePrefix(word string, ch byte) string {
	reverse := func(bs []byte, i, j int) {
		for i < j {
			bs[i], bs[j] = bs[j], bs[i]
			i++
			j--
		}
	}
	bs := []byte(word)
	i := 0
	for i < len(bs) && bs[i] != ch {
		i++
	}

	if i == len(bs) {
		return word
	}
	reverse(bs, 0, i)
	return string(bs)
}

// @lc code=end
