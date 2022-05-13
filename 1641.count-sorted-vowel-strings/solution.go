package countsortedvowelstrings

/*
 * @lc app=leetcode id=1641 lang=golang
 *
 * [1641] Count Sorted Vowel Strings
 */

// @lc code=start
func countVowelStrings(n int) int {
	var backtrack func(pool, n int) int
	backtrack = func(pool, n int) int {
		if n == 1 {
			return pool
		}
		cnt := 0
		for i := pool; i > 0; i-- {
			cnt += backtrack(i, n-1)
		}
		return cnt
	}

	return backtrack(5, n)
}

// @lc code=end
