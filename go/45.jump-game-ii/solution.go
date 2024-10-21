package jumpgameii

/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {
	steps := make([]int, len(nums))
	steps[0] = 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if j+nums[j] >= i {
				if steps[i] == 0 {
					steps[i] = steps[j] + 1
				} else {
					steps[i] = min(steps[i], steps[j]+1)
				}
			}
		}
	}

	return steps[len(nums)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
