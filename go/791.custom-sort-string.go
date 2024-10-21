package leetcodesolution

import "strings"

/*
 * @lc app=leetcode id=791 lang=golang
 *
 * [791] Custom Sort String
 */

// @lc code=start
func customSortString(order string, s string) string {
	m := [26]int{}
	for _, v := range s {
		m[v-'a']++
	}
	sb := strings.Builder{}
	fill := func(c rune, cnt int) {
		for i := 0; i < cnt; i++ {
			sb.WriteRune(c)
		}
	}
	for _, v := range order {
		fill(v, m[v-'a'])
		m[v-'a'] = 0
	}
	for i, v := range m {
		if v > 0 {
			fill(rune(i+'a'), v)
		}
	}

	return sb.String()
}

// @lc code=end
