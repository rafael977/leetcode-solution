package longestsubstringwithoutrepeatingcharacters

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) (ans int) {
	i := 0
	vi := make(map[byte]int)
	for j := range s {
		ci, visited := vi[s[j]]
		if visited && ci >= i {
			ans = max(ans, j-i)
			i = ci + 1
		}
		vi[s[j]] = j
	}
	ans = max(ans, len(s)-i)

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
