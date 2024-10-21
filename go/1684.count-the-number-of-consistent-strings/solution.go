package countthenumberofconsistentstrings

/*
 * @lc app=leetcode id=1684 lang=golang
 *
 * [1684] Count the Number of Consistent Strings
 */

// @lc code=start
func countConsistentStrings(allowed string, words []string) (ans int) {
	cm := make([]bool, 26)
	for _, v := range allowed {
		cm[v-'a'] = true
	}
	for _, w := range words {
		cnt := 1
		for _, c := range w {
			if !cm[c-'a'] {
				cnt = 0
				break
			}
		}
		ans += cnt
	}
	return
}

// @lc code=end
