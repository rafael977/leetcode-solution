package firstuniquecharacterinastring

/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */

// @lc code=start
func firstUniqChar(s string) (ans int) {
	// imap := make(map[rune]int)
	// had := make(map[rune]bool)
	// for i, c := range s {
	// 	if had[c] {
	// 		delete(imap, c)
	// 	} else {
	// 		imap[c] = i
	// 		had[c] = true
	// 	}
	// }

	// if len(imap) == 0 {
	// 	return -1
	// }
	// ans = len(s)
	// for _, i := range imap {
	// 	if ans > i {
	// 		ans = i
	// 	}
	// }
	// return
	c := [26]int{}
	for _, v := range s {
		c[v-'a']++
	}
	for i, v := range s {
		if c[v-'a'] == 1 {
			return i
		}
	}
	return -1
}

// @lc code=end
