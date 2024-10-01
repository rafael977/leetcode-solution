package checkifarraypairsaredivisiblebyk

/*
 * @lc app=leetcode id=1497 lang=golang
 *
 * [1497] Check If Array Pairs Are Divisible by k
 */

// @lc code=start
func canArrange(arr []int, k int) bool {
	rm := make(map[int]int)
	for _, v := range arr {
		r := v % k
		if r < 0 {
			r += k
		}
		rm[r]++
	}

	if rm[0]&1 != 0 {
		return false
	}
	if k&1 == 0 && rm[k/2]&1 != 0 {
		return false
	}
	delete(rm, 0)
	if k&1 == 0 {
		delete(rm, k/2)
	}
	for key, v := range rm {
		if rm[k-key] != v {
			return false
		}
	}

	return true
}

// @lc code=end
