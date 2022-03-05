package main

/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	sr, tr := []rune(s), []rune(t)
	si, ti := 0, 0

	for si < len(sr) {
		for ti < len(tr) && sr[si] != tr[ti] {
			ti++
		}
		if ti < len(tr) {
			si++
			ti++
		} else {
			break
		}
	}
	if si == len(sr) {
		return true
	} else {
		return false
	}
}

// @lc code=end
