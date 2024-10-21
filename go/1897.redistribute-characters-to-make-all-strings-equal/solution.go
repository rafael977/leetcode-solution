package redistributecharacterstomakeallstringsequal

/*
 * @lc app=leetcode id=1897 lang=golang
 *
 * [1897] Redistribute Characters to Make All Strings Equal
 */

// @lc code=start
func makeEqual(words []string) bool {
	cc := make([]int, 26)
	for _, s := range words {
		for _, c := range s {
			cc[c-'a']++
		}
	}

	n := len(words)
	for _, v := range cc {
		if v%n != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
