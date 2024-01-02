package convertanarrayintoa2darraywithconditions

/*
 * @lc app=leetcode id=2610 lang=golang
 *
 * [2610] Convert an Array Into a 2D Array With Conditions
 */

// @lc code=start
func findMatrix(nums []int) [][]int {
	rows := make([]map[int]bool, 0, len(nums))
	for _, v := range nums {
		inserted := false
		for i := range rows {
			if !rows[i][v] {
				rows[i][v] = true
				inserted = true
				break
			}
		}

		if !inserted {
			nr := make(map[int]bool)
			nr[v] = true
			rows = append(rows, nr)
		}
	}

	ans := make([][]int, 0, len(rows))
	for _, r := range rows {
		row := make([]int, 0, len(r))
		for k := range r {
			row = append(row, k)
		}

		ans = append(ans, row)
	}
	return ans
}

// @lc code=end
