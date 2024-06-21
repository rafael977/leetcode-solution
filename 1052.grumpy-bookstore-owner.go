package leetcodesolution

/*
 * @lc app=leetcode id=1052 lang=golang
 *
 * [1052] Grumpy Bookstore Owner
 */

// @lc code=start
func maxSatisfied(customers []int, grumpy []int, minutes int) (ans int) {
	s := 0
	for i, c := range customers {
		if grumpy[i] == 0 {
			s += c
		}
	}

	add := 0
	for i := 0; i < minutes; i++ {
		if grumpy[i] == 1 {
			add += customers[i]
		}
	}
	maxAdd := add
	for i := minutes; i < len(customers); i++ {
		if grumpy[i-minutes] == 1 {
			add -= customers[i-minutes]
		}
		if grumpy[i] == 1 {
			add += customers[i]
		}
		maxAdd = max(maxAdd, add)
	}

	return s + maxAdd
}

// @lc code=end
