package detonatethemaximumbombs

/*
 * @lc app=leetcode id=2101 lang=golang
 *
 * [2101] Detonate the Maximum Bombs
 */

// @lc code=start
func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	edges := make(map[int][]int)
	for i := 0; i < n; i++ {
		edges[i] = make([]int, 0)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			b1, b2 := bombs[i], bombs[j]
			x1, y1, r1 := b1[0], b1[1], b1[2]
			x2, y2, r2 := b2[0], b2[1], b2[2]
			d2 := (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
			if d2 <= r1*r1 {
				edges[i] = append(edges[i], j)
			}
			if d2 <= r2*r2 {
				edges[j] = append(edges[j], i)
			}
		}
	}

	max := 0
	for i := 0; i < n; i++ {
		visited := make([]bool, n)
		q := []int{i}
		cnt := 0
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			if visited[v] {
				continue
			}

			visited[v] = true
			cnt++
			for _, x := range edges[v] {
				if !visited[x] {
					q = append(q, x)
				}
			}
		}

		if max < cnt {
			max = cnt
		}
	}

	return max
}

// @lc code=end
