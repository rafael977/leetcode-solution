package findifpathexistsingraph

/*
 * @lc app=leetcode id=1971 lang=golang
 *
 * [1971] Find if Path Exists in Graph
 */

// @lc code=start
func validPath(n int, edges [][]int, source int, destination int) bool {
	parent := make(map[int]int)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if x == parent[x] {
			return x
		}
		p := find(parent[x])
		parent[x] = p
		return p
	}
	union := func(x, y int) {
		if x < y {
			parent[y] = x
		} else {
			parent[x] = y
		}
	}

	for _, e := range edges {
		p1, p2 := find(e[0]), find(e[1])
		if p1 != p2 {
			union(p1, p2)
		}
	}

	return find(source) == find(destination)
}

// @lc code=end
