package sumofprefixscoresofstrings

/*
 * @lc app=leetcode id=2416 lang=golang
 *
 * [2416] Sum of Prefix Scores of Strings
 */

// @lc code=start
func sumPrefixScores(words []string) []int {
	pfm := map[string]int{}
	for _, w := range words {
		for w != "" {
			pfm[w]++
			w = w[:len(w)-1]
		}
	}

	ans := make([]int, len(words))
	for i, w := range words {
		for w != "" {
			ans[i] += pfm[w]
			w = w[:len(w)-1]
		}
	}
	return ans
}

// @lc code=end
