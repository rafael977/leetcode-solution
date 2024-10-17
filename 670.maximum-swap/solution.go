package maximumswap

import "strconv"

/*
 * @lc app=leetcode id=670 lang=golang
 *
 * [670] Maximum Swap
 */

// @lc code=start
func maximumSwap(num int) int {
	sb := []byte(strconv.Itoa(num))
	for i := 0; i < len(sb)-1; i++ {
		j := i + 1
		for k := j; k < len(sb); k++ {
			if sb[j] <= sb[k] {
				j = k
			}
		}
		if sb[i] < sb[j] {
			sb[i], sb[j] = sb[j], sb[i]
			break
		}
	}

	n, _ := strconv.Atoi(string(sb))
	return n
}

// @lc code=end
