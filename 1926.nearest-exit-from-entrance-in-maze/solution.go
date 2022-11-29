package nearestexitfromentranceinmaze

/*
 * @lc app=leetcode id=1926 lang=golang
 *
 * [1926] Nearest Exit from Entrance in Maze
 */

// @lc code=start
var dirs [][]int = [][]int{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	q := [][]int{{entrance[0], entrance[1], 0}}
	maze[entrance[0]][entrance[1]] = '+'
	for len(q) > 0 {
		pos := q[0]
		q = q[1:]
		for _, d := range dirs {
			pp := []int{pos[0] + d[0], pos[1] + d[1]}
			if pp[0] < 0 || pp[0] >= m || pp[1] < 0 || pp[1] >= n ||
				maze[pp[0]][pp[1]] == '+' {
				continue
			}
			if pp[0] == 0 || pp[0] == m-1 || pp[1] == 0 || pp[1] == n-1 {
				return pos[2] + 1
			}
			q = append(q, []int{pp[0], pp[1], pos[2] + 1})
			maze[pp[0]][pp[1]] = '+'
		}
	}
	return -1
}

// @lc code=end
