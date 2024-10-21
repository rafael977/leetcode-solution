package poorpigs

import "math"

/*
 * @lc app=leetcode id=458 lang=golang
 *
 * [458] Poor Pigs
 */

// @lc code=start
func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	base := minutesToTest/minutesToDie + 1
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(base))))
}

// @lc code=end
