package leetcodesolution

/*
 * @lc app=leetcode id=1971 lang=golang
 *
 * [1971] Find if Path Exists in Graph
 */

// @lc code=start
// func validPath(n int, edges [][]int, source int, destination int) bool {
// 	parent := make(map[int]int)
// 	for i := 0; i < n; i++ {
// 		parent[i] = i
// 	}

// 	var find func(x int) int
// 	find = func(x int) int {
// 		if x == parent[x] {
// 			return x
// 		}
// 		p := find(parent[x])
// 		parent[x] = p
// 		return p
// 	}
// 	union := func(x, y int) {
// 		if x < y {
// 			parent[y] = x
// 		} else {
// 			parent[x] = y
// 		}
// 	}

// 	for _, e := range edges {
// 		p1, p2 := find(e[0]), find(e[1])
// 		if p1 != p2 {
// 			union(p1, p2)
// 		}
// 	}

// 	return find(source) == find(destination)
// }

func validPath(n int, edges [][]int, source int, destination int) bool {
	es := make(map[int][]int)
	for _, e := range edges {
		if _, ok := es[e[0]]; !ok {
			es[e[0]] = make([]int, 0)
		}
		if _, ok := es[e[1]]; !ok {
			es[e[1]] = make([]int, 0)
		}
		es[e[0]] = append(es[e[0]], e[1])
		es[e[1]] = append(es[e[1]], e[0])
	}

	visited := make(map[int]struct{}, n)
	q := []int{destination}
	for len(q) > 0 {
		v := q[len(q)-1]
		q = q[:len(q)-1]
		visited[v] = struct{}{}
		if v == source {
			return true
		}
		for _, n := range es[v] {
			if _, yes := visited[n]; yes {
				continue
			}
			q = append(q, n)
		}
	}
	return false
}

// @lc code=end
