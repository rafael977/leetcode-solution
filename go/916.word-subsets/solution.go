package wordsubsets

/*
 * @lc app=leetcode id=916 lang=golang
 *
 * [916] Word Subsets
 */

// @lc code=start
func wordSubsets(words1 []string, words2 []string) (ans []string) {
	count := func(w string) []int {
		m := make([]int, 26)
		for _, c := range w {
			m[c-'a']++
		}
		return m
	}

	w2c := make([]int, 26)
	for _, w := range words2 {
		cs := count(w)
		for i := range w2c {
			w2c[i] = max(w2c[i], cs[i])
		}
	}

	universal := func(w string) bool {
		wc := count(w)
		for i := range wc {
			if wc[i] < w2c[i] {
				return false
			}
		}
		return true
	}

	for _, w := range words1 {
		if universal(w) {
			ans = append(ans, w)
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
