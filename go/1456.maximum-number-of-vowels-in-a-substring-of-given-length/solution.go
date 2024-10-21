package maximumnumberofvowelsinasubstringofgivenlength

/*
 * @lc app=leetcode id=1456 lang=golang
 *
 * [1456] Maximum Number of Vowels in a Substring of Given Length
 */

// @lc code=start
func maxVowels(s string, k int) (ans int) {
	cnt := 0
	for _, v := range s[:k] {
		if v == 'a' ||
			v == 'e' ||
			v == 'i' ||
			v == 'o' ||
			v == 'u' {
			cnt++
		}
	}
	ans = cnt

	for i := k; i < len(s); i++ {
		if s[i] == 'a' ||
			s[i] == 'e' ||
			s[i] == 'i' ||
			s[i] == 'o' ||
			s[i] == 'u' {
			cnt++
		}
		if s[i-k] == 'a' ||
			s[i-k] == 'e' ||
			s[i-k] == 'i' ||
			s[i-k] == 'o' ||
			s[i-k] == 'u' {
			cnt--
		}

		if ans < cnt {
			ans = cnt
		}
	}
	return
}

// @lc code=end
