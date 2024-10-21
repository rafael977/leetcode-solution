package largestsubstringbetweentwoequalcharacters

/*
 * @lc app=leetcode id=1624 lang=golang
 *
 * [1624] Largest Substring Between Two Equal Characters
 */

// @lc code=start
func maxLengthBetweenEqualCharacters(s string) int {
	ans := -1
	idx := make([]int, 26)
	for i := range idx {
		idx[i] = -1
	}
	for i, c := range s {
		if idx[c-'a'] != -1 {
			ans = max(ans, i-idx[c-'a']-1)
		} else {
			idx[c-'a'] = i
		}
	}
	return ans
}

// @lc code=end
