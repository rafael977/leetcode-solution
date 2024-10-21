package leetcodesolution

/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	canMap := func(a, b string) bool {
		cm := make(map[byte]byte)
		for i := 0; i < len(a); i++ {
			c, ok := cm[a[i]]
			if !ok {
				cm[a[i]] = b[i]
				continue
			}
			if c != b[i] {
				return false
			}
		}
		return true
	}
	return canMap(s, t) && canMap(t, s)
}

// @lc code=end
