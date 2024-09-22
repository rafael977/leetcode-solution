package kthsmallestinlexicographicalorder

import (
	"math"
	"strconv"
)

/*
 * @lc app=leetcode id=440 lang=golang
 *
 * [440] K-th Smallest in Lexicographical Order
 */

// @lc code=start
func findKthNumber(n int, k int) int {
	getCnt := func(x, limit int) (c int) {
		k := len(strconv.Itoa(limit)) - len(strconv.Itoa(x))
		u := limit
		for i := 0; i < k; i++ {
			c += int(math.Pow10(i))
			u /= 10
		}
		if u > x {
			c += int(math.Pow10(k))
		} else if u == x {
			c += limit - x*int(math.Pow10(k)) + 1
		}
		return
	}

	x := 1
	for k > 1 {
		cnt := getCnt(x, n)
		if cnt < k {
			x++
			k -= cnt
		} else {
			x *= 10
			k -= 1
		}
	}

	return x
}

// @lc code=end
