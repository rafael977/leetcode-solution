package divideplayersintoteamsofequalskill

import "sort"

/*
 * @lc app=leetcode id=2491 lang=golang
 *
 * [2491] Divide Players Into Teams of Equal Skill
 */

// @lc code=start
func dividePlayers(skill []int) (ans int64) {
	n := len(skill)
	sort.Ints(skill)
	ts := skill[0] + skill[n-1]
	for i := 0; i < n/2; i++ {
		if skill[i]+skill[n-1-i] != ts {
			return -1
		}
		ans += int64(skill[i] * skill[n-1-i])
	}
	return
}

// @lc code=end
