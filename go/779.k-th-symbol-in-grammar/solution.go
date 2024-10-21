package kthsymbolingrammar

/*
 * @lc app=leetcode id=779 lang=golang
 *
 * [779] K-th Symbol in Grammar
 */

// @lc code=start
func kthGrammar(n int, k int) int {
	if n == 1 && k == 1 {
		return 0
	}

	pre := kthGrammar(n-1, (k+1)/2)
	if k%2 == 1 {
		return pre
	} else {
		return (pre + 1) % 2
	}
}

// @lc code=end
