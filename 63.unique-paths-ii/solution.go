package uniquepathsii

/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	paths := make([][]int, m)
	for i := range paths {
		paths[i] = make([]int, n)
	}
	for i := range paths {
		if obstacleGrid[i][0] == 1 {
			break
		}
		paths[i][0] = 1
	}
	for i := range paths[0] {
		if obstacleGrid[0][i] == 1 {
			break
		}
		paths[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				paths[i][j] = paths[i-1][j] + paths[i][j-1]
			}
		}
	}

	return paths[m-1][n-1]
}

// @lc code=end
