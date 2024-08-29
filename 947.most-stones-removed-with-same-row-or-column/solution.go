package moststonesremovedwithsameroworcolumn

/*
 * @lc app=leetcode id=947 lang=golang
 *
 * [947] Most Stones Removed with Same Row or Column
 */

// @lc code=start
func removeStones(stones [][]int) (ans int) {
	n := len(stones)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}

	var find func(i int) int
	find = func(i int) int {
		if i != p[i] {
			p[i] = find(p[i])
		}
		return p[i]
	}
	union := func(x, y int) {
		px, py := find(x), find(y)
		p[py] = px
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if stones[i][0] == stones[j][0] ||
				stones[i][1] == stones[j][1] {
				union(i, j)
			}
		}
	}

	for i := range p {
		if p[i] != i {
			ans++
		}
	}

	return
}

// @lc code=end
