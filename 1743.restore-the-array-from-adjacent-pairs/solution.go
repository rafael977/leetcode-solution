package restorethearrayfromadjacentpairsgo

/*
 * @lc app=leetcode id=1743 lang=golang
 *
 * [1743] Restore the Array From Adjacent Pairs
 */

// @lc code=start
func restoreArray(adjacentPairs [][]int) (ans []int) {
	link := make(map[int][]int)
	for _, v := range adjacentPairs {
		a, b := v[0], v[1]
		if link[a] == nil {
			link[a] = make([]int, 0, 2)
		}
		if link[b] == nil {
			link[b] = make([]int, 0, 2)
		}
		link[a] = append(link[a], b)
		link[b] = append(link[b], a)
	}

	vi := make(map[int]bool)
	for k, v := range link {
		if len(v) == 1 {
			ans = append(ans, k)
			vi[k] = true
			break
		}
	}

	for i := 0; i < len(link)-1; i++ {
		p := ans[i]
		for _, v := range link[p] {
			if vi[v] {
				continue
			}
			ans = append(ans, v)
			vi[v] = true
		}
	}
	return
}

// @lc code=end
