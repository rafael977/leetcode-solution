package minimumroundstocompletealltasks

/*
 * @lc app=leetcode id=2244 lang=golang
 *
 * [2244] Minimum Rounds to Complete All Tasks
 */

// @lc code=start
func minimumRounds(tasks []int) (ans int) {
	tc := make(map[int]int)
	for _, v := range tasks {
		tc[v]++
	}

	for _, v := range tc {
		if v == 1 {
			return -1
		}
		if v%3 == 0 {
			ans += v / 3
		} else {
			ans += v/3 + 1
		}
	}
	return
}

// @lc code=end
