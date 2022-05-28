package numberof1bits

/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

// @lc code=start
func hammingWeight(num uint32) (d int) {
	for num > 0 {
		d += int(num & 1)
		num = num >> 1
	}
	return
}

// @lc code=end
