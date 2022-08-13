package substringwithconcatenationofallwords

/*
 * @lc app=leetcode id=30 lang=golang
 *
 * [30] Substring with Concatenation of All Words
 */

// @lc code=start
func findSubstring(s string, words []string) (ans []int) {
	for i := 0; i < len(words[0]); i++ {
		ans = append(ans, find(s[i:], i, words)...)
	}
	return
}

func find(s string, si int, words []string) (ans []int) {
	lw, n := len(words), len(words[0])
	if len(s) < lw*n {
		return
	}

	diff := make(map[string]int)
	for i := 0; i < lw*n; i += n {
		diff[s[i:i+n]]++
	}
	for _, w := range words {
		diff[w]--
		if diff[w] == 0 {
			delete(diff, w)
		}
	}

	for i := 0; i <= len(s)/n-lw; i++ {
		if i != 0 {
			w := s[(i-1)*n : i*n]
			diff[w]--
			if diff[w] == 0 {
				delete(diff, w)
			}
			w = s[(i+lw-1)*n : (i+lw)*n]
			diff[w]++
			if diff[w] == 0 {
				delete(diff, w)
			}
		}
		if len(diff) == 0 {
			ans = append(ans, si+i*n)
		}
	}
	return
}

// @lc code=end
