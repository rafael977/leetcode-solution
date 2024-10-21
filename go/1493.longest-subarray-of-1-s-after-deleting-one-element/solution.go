package longestsubarrayof1safterdeletingoneelement

/*
 * @lc app=leetcode id=1493 lang=golang
 *
 * [1493] Longest Subarray of 1's After Deleting One Element
 */

// @lc code=start
func longestSubarray(nums []int) int {
	cnts := make([]int, 0)
	cc := 0
	for _, v := range nums {
		if v == 0 {
			cnts = append(cnts, cc)
			cc = 0
			continue
		}
		cc++
	}
	if cc > 0 {
		cnts = append(cnts, cc)
	}

	if len(cnts) == 0 {
		return 0
	} else if len(cnts) == 1 {
		return cnts[0] - 1
	}

	max := 0
	for i := 1; i < len(cnts); i++ {
		s := cnts[i-1] + cnts[i]
		if s > max {
			max = s
		}
	}
	return max
}

// @lc code=end
