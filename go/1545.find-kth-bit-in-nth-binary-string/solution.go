package findkthbitinnthbinarystring

import "math"

/*
 * @lc app=leetcode id=1545 lang=golang
 *
 * [1545] Find Kth Bit in Nth Binary String
 */

// @lc code=start
func findKthBit(n int, k int) byte {
	if n == 1 {
		return '0'
	}
	l := int(math.Pow(2, float64(n))) - 1
	half := l / 2
	if k == half+1 {
		return '1'
	} else if k <= half {
		return findKthBit(n-1, k)
	} else {
		p := half + 1 - (k - half - 1)
		return revert(findKthBit(n-1, p))
	}
}

func revert(x byte) byte {
	if x == '0' {
		return '1'
	}
	return '0'
}

// @lc code=end
