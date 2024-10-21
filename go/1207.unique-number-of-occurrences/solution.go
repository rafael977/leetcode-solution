package uniquenumberofoccurrences

/*
 * @lc app=leetcode id=1207 lang=golang
 *
 * [1207] Unique Number of Occurrences
 */

// @lc code=start
func uniqueOccurrences(arr []int) bool {
    counts := make(map[int]int)
	for _, v := range arr {
		counts[v]++
	}

	occur := make(map[int]bool)
	for _, v := range counts {
		if _, ok := occur[v]; ok {
			return false
		}
		occur[v] = true
	}
	return true
}
// @lc code=end

