package pascalstriangle

/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */

// @lc code=start
var t [][]int

func generate(numRows int) [][]int {
	if len(t) == 0 {
		t = append(t, []int{1})
	}
	if numRows <= len(t) {
		return t[:numRows]
	}
	for i := len(t); i < numRows; i++ {
		j := i - 1
		row := make([]int, len(t[j])+1)
		for k := range row {
			if k == 0 || k == len(row)-1 {
				row[k] = 1
			} else {
				row[k] = t[j][k-1] + t[j][k]
			}
		}
		t = append(t, row)
	}
	return t
}

// @lc code=end
