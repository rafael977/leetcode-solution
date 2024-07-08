package leetcodesolution

/*
 * @lc app=leetcode id=1823 lang=golang
 *
 * [1823] Find the Winner of the Circular Game
 */

// @lc code=start
func findTheWinner(n int, k int) int {
	players := make([]int, n)
	for i := range players {
		players[i] = i + 1
	}

	i := 0
	for len(players) > 1 {
		i = (i + k - 1) % len(players)
		players = append(players[:i], players[i+1:]...)
	}
	return players[0]
}

// @lc code=end
