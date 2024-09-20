package shortestpalindrome

import "fmt"

/*
 * @lc app=leetcode id=214 lang=golang
 *
 * [214] Shortest Palindrome
 */

// @lc code=start
func shortestPalindrome(s string) string {
	rs := reverse(s)
	p := fmt.Sprintf("%s#%s", s, rs)
	next := computeNext(p)

	return rs[:len(rs)-next[len(next)-1]] + s
}

func computeNext(pattern string) []int {
	next := make([]int, len(pattern)+1)
	next[0] = -1
	next[1] = 0
	i, k := 2, 0
	for i <= len(pattern) {
		if pattern[i-1] == pattern[k] {
			k++
			next[i] = k
			i++
		} else if k == 0 {
			next[i] = 0
			i++
		} else {
			k = next[k]
		}
	}
	return next
}

func reverse(s string) string {
	bs := []byte(s)
	for i := 0; i < len(bs)/2; i++ {
		bs[i], bs[len(s)-1-i] = bs[len(s)-1-i], bs[i]
	}
	return string(bs)
}

// @lc code=end
