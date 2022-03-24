package main

import "sort"

/*
 * @lc app=leetcode id=881 lang=golang
 *
 * [881] Boats to Save People
 */

// @lc code=start
func numRescueBoats(people []int, limit int) (ans int) {
	sort.Ints(people)
	i, j := 0, len(people)-1
	for i < j {
		if people[i]+people[j] <= limit {
			i++
			j--
			ans++
		} else {
			j--
			ans++
		}
	}
	if i == j {
		ans++
	}
	return
}

// @lc code=end
