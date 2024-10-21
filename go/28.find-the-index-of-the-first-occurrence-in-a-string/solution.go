package findtheindexofthefirstoccurrenceinastring

/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Find the Index of the First Occurrence in a String
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	for i := range haystack {
		j := 0
		for ; j < len(needle) && i+j < len(haystack); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}

// @lc code=end
