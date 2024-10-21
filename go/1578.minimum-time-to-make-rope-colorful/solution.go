package minimumtimetomakeropecolorful

/*
 * @lc app=leetcode id=1578 lang=golang
 *
 * [1578] Minimum Time to Make Rope Colorful
 */

// @lc code=start
func minCost(colors string, neededTime []int) (ans int) {
	i := 0
	for i < len(colors) {
		ch := colors[i]
		sum, max := 0, 0
		for i < len(colors) && ch == colors[i] {
			sum += neededTime[i]
			if max < neededTime[i] {
				max = neededTime[i]
			}
			i++
		}
		ans += sum - max
	}
	return
}

// @lc code=end
