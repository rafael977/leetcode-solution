package leetcodesolution

/*
 * @lc app=leetcode id=860 lang=golang
 *
 * [860] Lemonade Change
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	cash := []int{0, 0}
	for _, b := range bills {
		switch b {
		case 5:
			cash[0]++
		case 10:
			if cash[0] > 0 {
				cash[0]--
				cash[1]++
			} else {
				return false
			}
		case 20:
			if cash[0] > 0 && cash[1] > 0 {
				cash[0]--
				cash[1]--
			} else if cash[0] >= 3 {
				cash[0] -= 3
			} else {
				return false
			}
		}
	}

	return true
}

// @lc code=end
