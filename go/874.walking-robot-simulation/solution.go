package walkingrobotsimulation

/*
 * @lc app=leetcode id=874 lang=golang
 *
 * [874] Walking Robot Simulation
 */

// @lc code=start
func robotSim(commands []int, obstacles [][]int) (ans int) {
	op := make(map[int]map[int]bool)
	for _, v := range obstacles {
		if _, ok := op[v[0]]; !ok {
			op[v[0]] = make(map[int]bool)
		}
		op[v[0]][v[1]] = true
	}

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	p := [2]int{0, 0}
	d := 0
	for _, c := range commands {
		if c > 0 {
			dir := dirs[d]
			for i := 0; i < c; i++ {
				np := [2]int{p[0] + dir[0], p[1] + dir[1]}
				if op[np[0]][np[1]] {
					break
				}
				p = np
				ans = max(ans, p[0]*p[0]+p[1]*p[1])
			}
			continue
		}
		if c == -1 {
			d++
		}
		if c == -2 {
			d--
		}
		d = (d%4 + 4) % 4
	}

	return
}

// @lc code=end
