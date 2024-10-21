package groupthepeoplegiventhegroupsizetheybelongto

import "sort"

/*
 * @lc app=leetcode id=1282 lang=golang
 *
 * [1282] Group the People Given the Group Size They Belong To
 */

// @lc code=start
func groupThePeople(groupSizes []int) (ans [][]int) {
	idxG := make([][2]int, 0)
	for i, v := range groupSizes {
		idxG = append(idxG, [2]int{v, i})
	}

	sort.Slice(idxG, func(i, j int) bool {
		if idxG[i][0] == idxG[j][0] {
			return idxG[i][1] < idxG[j][1]
		}
		return idxG[i][0] < idxG[j][0]
	})

	g := make([]int, 0)
	for _, v := range idxG {
		g = append(g, v[1])
		if len(g) == v[0] {
			ans = append(ans, append([]int{}, g...))
			g = make([]int, 0)
		}
	}
	return
}

// @lc code=end
