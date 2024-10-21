package validanagram

/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	dic := make([]int, 26)
	for _, c := range s {
		dic[c-'a']++
	}
	for _, c := range t {
		dic[c-'a']--
	}

	for _, c := range dic {
		if c != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
