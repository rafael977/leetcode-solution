package maximumwidthramp

/*
 * @lc app=leetcode id=962 lang=golang
 *
 * [962] Maximum Width Ramp
 */

// @lc code=start
func maxWidthRamp(nums []int) (ans int) {
	n := len(nums)
	st := make([]int, 0, n)
	for i, v := range nums {
		if len(st) == 0 || nums[st[len(st)-1]] > v {
			st = append(st, i)
		}
	}

	for j := n - 1; j >= 0; j-- {
		for len(st) > 0 && nums[j] >= nums[st[len(st)-1]] {
			ans = max(ans, j-st[len(st)-1])
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			break
		}
	}
	return
}

// @lc code=end
