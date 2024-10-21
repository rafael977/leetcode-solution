package leetcodesolution

/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */

// @lc code=start
func trap(height []int) (ans int) {
	st := make([]int, 0, len(height))
	for i, v := range height {
		for len(st) > 0 && v > height[st[len(st)-1]] {
			top := st[len(st)-1]
			st = st[:len(st)-1]
			if len(st) == 0 {
				break
			}
			left := st[len(st)-1]
			ans += (i - left - 1) * (min(height[left], v) - height[top])
		}

		st = append(st, i)
	}
	return
}

// @lc code=end
