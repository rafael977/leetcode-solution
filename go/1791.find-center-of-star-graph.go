package leetcodesolution

/*
 * @lc app=leetcode id=1791 lang=golang
 *
 * [1791] Find Center of Star Graph
 */

// @lc code=start
func findCenter(edges [][]int) int {
	ec := make(map[int]int)
	for _, e := range edges {
		ec[e[0]]++
		ec[e[1]]++
	}

	for x, v := range ec {
		if v == len(ec)-1 {
			return x
		}
	}
	return 0
}

// @lc code=end
