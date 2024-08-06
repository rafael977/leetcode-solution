package leetcodesolution

import "sort"

/*
 * @lc app=leetcode id=3016 lang=golang
 *
 * [3016] Minimum Number of Pushes to Type Word II
 */

// @lc code=start
func minimumPushes(word string) (ans int) {
	c := make([]int, 26)
	for _, v := range word {
		c[v-'a']++
	}
	sort.Ints(c)
	for i := 25; i >= 0; i-- {
		if c[i] == 0 {
			break
		}
		ans += c[i] * ((25-i)/8 + 1)
	}
	return
}

// @lc code=end
