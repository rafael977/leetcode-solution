package findkclosestelements

import "sort"

/*
 * @lc app=leetcode id=658 lang=golang
 *
 * [658] Find K Closest Elements
 */

// @lc code=start
func findClosestElements(arr []int, k int, x int) (ans []int) {
	j := sort.SearchInts(arr, x)
	i := j - 1
	for len(ans) < k {
		if i < 0 {
			ans = append(ans, arr[j])
			j++
		} else if j >= len(arr) {
			ans = append(ans, arr[i])
			i--
		} else {
			if closer(arr[i], arr[j], x) {
				ans = append(ans, arr[i])
				i--
			} else {
				ans = append(ans, arr[j])
				j++
			}
		}
	}
	sort.Ints(ans)
	return
}

func closer(a, b, x int) bool {
	da, db := a-x, b-x
	if da < 0 {
		da *= -1
	}
	if db < 0 {
		db *= -1
	}

	if da == db {
		return a < b
	}
	return da < db
}

// @lc code=end
