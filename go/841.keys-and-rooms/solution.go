package keysandrooms

/*
 * @lc app=leetcode id=841 lang=golang
 *
 * [841] Keys and Rooms
 */

// @lc code=start
func canVisitAllRooms(rooms [][]int) bool {
	parent := make([]int, len(rooms))
	for i := range parent {
		parent[i] = i
	}

	var find func(i int) int
	find = func(i int) int {
		if i == parent[i] {
			return i
		}
		parent[i] = find(parent[i])
		return parent[i]
	}
	union := func(i, j int) {
		if i < j {
			parent[j] = i
		} else {
			parent[i] = j
		}
	}

	q := []int{0}
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		for _, k := range rooms[i] {
			p1, p2 := find(i), find(k)
			if p1 != p2 {
				union(p1, p2)
				q = append(q, k)
			}
		}
	}

	for _, p := range parent {
		if p != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
