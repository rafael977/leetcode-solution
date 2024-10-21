package minimumstringlengthafterremovingsubstrings

/*
 * @lc app=leetcode id=2696 lang=golang
 *
 * [2696] Minimum String Length After Removing Substrings
 */

// @lc code=start
func minLength(s string) int {
	st := make([]rune, 0, len(s))
	for _, v := range s {
		if len(st) > 0 && (v == 'B' && st[len(st)-1] == 'A' ||
			(v == 'D' && st[len(st)-1] == 'C')) {
			st = st[:len(st)-1]
			continue
		}
		st = append(st, v)
	}
	return len(st)
}

// @lc code=end
