package sequentialdigits

import "strconv"

/*
 * @lc app=leetcode id=1291 lang=golang
 *
 * [1291] Sequential Digits
 */

// @lc code=start
var all []int

func sequentialDigits(low int, high int) []int {
	if len(all) == 0 {
		genAll()
	}

	lo, hi := 0, len(all)-1
	for lo < len(all) && all[lo] < low {
		lo++
	}
	for hi >= 0 && all[hi] > high {
		hi--
	}
	if lo > hi {
		return []int{}
	}
	return all[lo : hi+1]
}

func genAll() {
	s := "123456789"
	for d := 2; d <= 9; d++ {
		for i := 0; i <= 9-d; i++ {
			v, _ := strconv.Atoi(s[i : i+d])
			all = append(all, v)
		}
	}
}

// @lc code=end
