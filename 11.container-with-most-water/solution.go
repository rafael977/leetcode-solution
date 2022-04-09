package containerwithmostwater

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) (max int) {
	for i, j := 0, len(height)-1; i < j; {
		low := 0
		if height[i] < height[j] {
			low = height[i]
			i++
		} else {
			low = height[j]
			j--
		}

		area := low * (j - i + 1)
		if max < area {
			max = area
		}
	}
	return
}

// @lc code=end
