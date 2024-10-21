package maximumerasurevalue

/*
 * @lc app=leetcode id=1695 lang=golang
 *
 * [1695] Maximum Erasure Value
 */

// @lc code=start
func maximumUniqueSubarray(nums []int) (ans int) {
	vi := make(map[int]int)
	sum := 0
	i := 0
	for j := range nums {
		idx, had := vi[nums[j]]
		if had {
			for i < idx+1 {
				sum -= nums[i]
				i++
			}
		}
		vi[nums[j]] = j
		sum += nums[j]
		ans = max(ans, sum)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
