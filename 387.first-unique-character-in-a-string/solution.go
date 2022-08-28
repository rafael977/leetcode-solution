package firstuniquecharacterinastring

/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */

// @lc code=start
func firstUniqChar(s string) (ans int) {
	imap := make(map[rune]int)
	had := make(map[rune]bool)
	for i, c := range s {
		if had[c] {
			delete(imap, c)
		} else {
			imap[c] = i
			had[c] = true
		}
	}

	if len(imap) == 0 {
		return -1
	}
	ans = len(s)
	for _, i := range imap {
		if ans > i {
			ans = i
		}
	}
	return
}

// @lc code=end
