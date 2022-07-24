package numberofmatchingsubsequences

/*
 * @lc app=leetcode id=792 lang=golang
 *
 * [792] Number of Matching Subsequences
 */

// @lc code=start
func numMatchingSubseq(s string, words []string) (ans int) {
	heads := make([][]string, 26)
	for i := range heads {
		heads[i] = make([]string, 0)
	}
	for _, w := range words {
		heads[w[0]-'a'] = append(heads[w[0]-'a'], w)
	}

	for _, c := range []byte(s) {
		ws := heads[c-'a']
		heads[c-'a'] = make([]string, 0)
		for _, w := range ws {
			if len(w) == 1 {
				ans++
				continue
			}
			w = w[1:]
			heads[w[0]-'a'] = append(heads[w[0]-'a'], w)
		}
	}
	return
}

// @lc code=end
