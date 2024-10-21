package findandreplacepattern

/*
 * @lc app=leetcode id=890 lang=golang
 *
 * [890] Find and Replace Pattern
 */

// @lc code=start
func findAndReplacePattern(words []string, pattern string) (ans []string) {
	match := func(w, p string) bool {
		m, rm := make(map[byte]byte), make(map[byte]byte)
		for i := range w {
			mc, ok1 := m[w[i]]
			rmc, ok2 := rm[p[i]]
			if !ok1 && !ok2 {
				m[w[i]] = p[i]
				rm[p[i]] = w[i]
			} else if mc != p[i] || rmc != w[i] {
				return false
			}
		}
		return true
	}

	for _, w := range words {
		if len(w) != len(pattern) {
			continue
		}
		if match(w, pattern) {
			ans = append(ans, w)
		}
	}
	return
}

// @lc code=end
