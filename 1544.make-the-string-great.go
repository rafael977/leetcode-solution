package leetcodesolution

/*
 * @lc app=leetcode id=1544 lang=golang
 *
 * [1544] Make The String Great
 */

// @lc code=start
func makeGood(s string) string {
	st := make([]rune, 0, len(s))
	for _, v := range s {
		if len(st) > 0 && sameCharDiffCase(st[len(st)-1], v) {
			st = st[:len(st)-1]
			continue
		}
		st = append(st, v)
	}
	return string(st)
}

func sameCharDiffCase(a, b rune) bool {
	if a > b {
		return a-b == 'a'-'A'
	}
	return b-a == 'a'-'A'
}

// @lc code=end
