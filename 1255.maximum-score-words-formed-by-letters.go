package leetcodesolution

/*
 * @lc app=leetcode id=1255 lang=golang
 *
 * [1255] Maximum Score Words Formed by Letters
 */

// @lc code=start
func maxScoreWords(words []string, letters []byte, score []int) (ans int) {
	lc := make(map[byte]int)
	for _, v := range letters {
		lc[v]++
	}

	wlc := make(map[byte]int)
	var backtrack func(i int)
	backtrack = func(i int) {
		if i == len(words) {
			ok := true
			sc := 0
			for l, c := range wlc {
				if c > lc[l] {
					ok = false
					break
				}
				sc += c * score[l-'a']
			}

			if ok {
				ans = max(ans, sc)
			}
			return
		}
		for _, c := range []byte(words[i]) {
			wlc[c]++
		}
		backtrack(i + 1)
		for _, c := range []byte(words[i]) {
			wlc[c]--
		}
		backtrack(i + 1)
	}

	backtrack(0)
	return
}

// @lc code=end
