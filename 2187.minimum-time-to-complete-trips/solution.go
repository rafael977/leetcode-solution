package minimumtimetocompletetrips

import "sort"

/*
 * @lc app=leetcode id=2187 lang=golang
 *
 * [2187] Minimum Time to Complete Trips
 */

// @lc code=start
func minimumTime(time []int, totalTrips int) int64 {
	maxTime := 0
	for _, v := range time {
		if v > maxTime {
			maxTime = v
		}
	}

	t := sort.Search(totalTrips*maxTime+1, func(i int) bool {
		cnt := 0
		for _, v := range time {
			cnt += i / v
		}
		return cnt >= totalTrips
	})

	return int64(t)
}

// @lc code=end
