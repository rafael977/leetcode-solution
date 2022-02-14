package main

/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	c1, c2 := [26]int{}, [26]int{}
	for i := 0; i < n; i++ {
		c1[s1[i]-'a']++
		c2[s2[i]-'a']++
	}
	if c1 == c2 {
		return true
	}
	for i := n; i < m; i++ {
		c2[s2[i]-'a']++
		c2[s2[i-n]-'a']--
		if c1 == c2 {
			return true
		}
	}

	return false
}

// @lc code=end
