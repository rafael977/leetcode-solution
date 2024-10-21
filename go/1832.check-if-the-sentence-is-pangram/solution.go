package checkifthesentenceispangram

/*
 * @lc app=leetcode id=1832 lang=golang
 *
 * [1832] Check if the Sentence Is Pangram
 */

// @lc code=start
func checkIfPangram(sentence string) (ans bool) {
	has := [26]bool{}
	for _, v := range sentence {
		has[v-'a'] = true
	}

	ans = true
	for _, v := range has {
		ans = ans && v
	}
	return
}

// @lc code=end
