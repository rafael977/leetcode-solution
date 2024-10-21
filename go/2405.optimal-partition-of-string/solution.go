package optimalpartitionofstring

/*
 * @lc app=leetcode id=2405 lang=golang
 *
 * [2405] Optimal Partition of String
 */

// @lc code=start
func partitionString(s string) (ans int) {
	cc := make([]bool, 26)
	for _, v := range s {
		if !cc[v-'a'] {
			cc[v-'a'] = true
		} else {
			cc = make([]bool, 26)
			cc[v-'a'] = true
			ans++
		}
	}
	ans++
	return
}

// @lc code=end
