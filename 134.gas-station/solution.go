package gasstation

/*
 * @lc app=leetcode id=134 lang=golang
 *
 * [134] Gas Station
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) (ans int) {
	tank := 0
	less := 0
	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			if i == len(gas)-1 {
				ans = -1
			} else {
				ans = i + 1
				less += tank
				tank = 0
			}
		}
	}

	if ans != -1 && tank+less < 0 {
		return -1
	}
	return
}

// @lc code=end
