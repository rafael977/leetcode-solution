package maximumdistanceinarrays

/*
 * @lc app=leetcode id=624 lang=golang
 *
 * [624] Maximum Distance in Arrays
 */

// @lc code=start
func maxDistance(arrays [][]int) (ans int) {
	minIdx := 0
	for i, v := range arrays {
		if v[0] < arrays[minIdx][0] {
			minIdx = i
		}
	}

	minMin, minMax := arrays[minIdx][0], arrays[minIdx][len(arrays[minIdx])-1]
	for i := range arrays {
		if i == minIdx {
			continue
		}

		ans = max(ans, arrays[i][len(arrays[i])-1]-minMin)
		ans = max(ans, minMax-arrays[i][0])
	}

	return
}

// @lc code=end
