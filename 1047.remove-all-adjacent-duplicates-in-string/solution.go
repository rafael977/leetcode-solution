package removealladjacentduplicatesinstring

/*
 * @lc app=leetcode id=1047 lang=golang
 *
 * [1047] Remove All Adjacent Duplicates In String
 */

// @lc code=start
func removeDuplicates(s string) string {
	st := make([]byte, 0)
	for _, v := range []byte(s) {
		if len(st) == 0 || v != st[len(st)-1] {
			st = append(st, v)
		} else {
			st = st[:len(st)-1]
		}
	}

	return string(st)
}

// @lc code=end
