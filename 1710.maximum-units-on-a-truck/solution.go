package maximumunitsonatruck

import "sort"

/*
 * @lc app=leetcode id=1710 lang=golang
 *
 * [1710] Maximum Units on a Truck
 */

// @lc code=start
func maximumUnits(boxTypes [][]int, truckSize int) (ans int) {
	sort.Slice(boxTypes, func(i, j int) bool {
		a, b := boxTypes[i], boxTypes[j]
		return a[1] > b[1]
	})

	for _, b := range boxTypes {
		if b[0] > truckSize {
			ans += truckSize * b[1]
			break
		}
		ans += b[0] * b[1]
		truckSize -= b[0]
	}
	return
}

// @lc code=end
