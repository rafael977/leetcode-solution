package main

/*
 * @lc app=leetcode id=763 lang=golang
 *
 * [763] Partition Labels
 */

// @lc code=start
func partitionLabels(s string) (ans []int) {
	ls := [26]int{}
	for i, v := range s {
		ls[v-'a'] = i
	}

	start, end := 0, 0
	for i, v := range s {
		if ls[v-'a'] > end {
			end = ls[v-'a']
		}
		if i == end {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}
	return
}

// @lc code=end
