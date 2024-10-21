package leetcodesolution

/*
 * @lc app=leetcode id=840 lang=golang
 *
 * [840] Magic Squares In Grid
 */

// @lc code=start
func numMagicSquaresInside(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	if m < 3 || n < 3 {
		return 0
	}

	magic := func(v ...int) bool {
		count := make([]int, 9)
		for _, n := range v {
			if n-1 < 9 && n-1 >= 0 {
				count[n-1]++
			}
		}
		for _, c := range count {
			if c != 1 {
				return false
			}
		}

		return v[0]+v[1]+v[2] == 15 &&
			v[3]+v[4]+v[5] == 15 &&
			v[6]+v[7]+v[8] == 15 &&
			v[0]+v[4]+v[8] == 15 &&
			v[2]+v[4]+v[6] == 15 &&
			v[0]+v[3]+v[6] == 15 &&
			v[1]+v[4]+v[7] == 15 &&
			v[2]+v[5]+v[8] == 15
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] != 5 {
				continue
			}
			if magic(grid[i-1][j-1], grid[i-1][j], grid[i-1][j+1],
				grid[i][j-1], grid[i][j], grid[i][j+1],
				grid[i+1][j-1], grid[i+1][j], grid[i+1][j+1]) {
				ans++
			}
		}
	}

	return
}

// @lc code=end
