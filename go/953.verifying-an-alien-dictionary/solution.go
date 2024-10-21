package verifyinganaliendictionary

/*
 * @lc app=leetcode id=953 lang=golang
 *
 * [953] Verifying an Alien Dictionary
 */

// @lc code=start
func isAlienSorted(words []string, order string) bool {
	w := make(map[byte]int)
	for i := range order {
		w[order[i]] = i
	}

	for i := 1; i < len(words); i++ {
		if !isLess(words[i-1], words[i], w) {
			return false
		}
	}
	return true
}

func isLess(s1, s2 string, w map[byte]int) bool {
	i := 0
	for ; i < len(s1) && i < len(s2); i++ {
		if w[s1[i]] < w[s2[i]] {
			return true
		} else if w[s1[i]] > w[s2[i]] {
			return false
		}
	}

	return i >= len(s1)
}

// @lc code=end
