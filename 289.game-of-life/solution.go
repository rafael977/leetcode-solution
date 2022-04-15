package gameoflife

/*
 * @lc app=leetcode id=289 lang=golang
 *
 * [289] Game of Life
 */

// @lc code=start
func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			neighbors := 0
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if i == 0 && j == 0 {
						continue
					}
					r, c := row+i, col+j
					if (r >= 0 && r < m) && (c >= 0 && c < n) && board[r][c] > 0 {
						neighbors++
					}
				}
			}

			// rule 1 & 3, live -> dead, mark as 2
			if board[row][col] == 1 && (neighbors < 2 || neighbors > 3) {
				board[row][col] = 2
			}

			// rule 4, dead -> live, mark as -1
			if board[row][col] == 0 && neighbors == 3 {
				board[row][col] = -1
			}
		}
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if board[row][col] == 2 {
				board[row][col] = 0
			}
			if board[row][col] == -1 {
				board[row][col] = 1
			}
		}
	}
}

// @lc code=end
