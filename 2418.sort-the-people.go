package leetcodesolution

import "sort"

/*
 * @lc app=leetcode id=2418 lang=golang
 *
 * [2418] Sort the People
 */

// @lc code=start
func sortPeople(names []string, heights []int) (ans []string) {
	nh := make([]struct {
		n string
		h int
	}, len(names))
	for i, v := range names {
		nh[i] = struct {
			n string
			h int
		}{
			n: v,
			h: heights[i],
		}
	}
	sort.Slice(nh, func(i, j int) bool {
		return nh[i].h > nh[j].h
	})

	for _, v := range nh {
		ans = append(ans, v.n)
	}
	return
}

// @lc code=end
