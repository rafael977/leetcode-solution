package main

/*
 * @lc app=leetcode id=3713 lang=golang
 *
 * [3713] Longest Balanced Substring I
 */

// @lc code=start
func longestBalanced(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		cnt := [26]int{}
		for j := i; j < len(s); j++ {
			cnt[s[j]-'a']++
			x := cnt[s[j]-'a']
			isBalanced := true
			for _, y := range cnt {
				if y != 0 && y != x {
					isBalanced = false
					break
				}
			}
			if isBalanced {
				res = max(res, j-i+1)
			}
		}
	}
	return res
}

// @lc code=end
