package numberofnodesinthesubtreewiththesamelabel

/*
 * @lc app=leetcode id=1519 lang=golang
 *
 * [1519] Number of Nodes in the Sub-Tree With the Same Label
 */

// @lc code=start
func countSubTrees(n int, edges [][]int, labels string) (ans []int) {
	ans = make([]int, n)
	em := make(map[int][]int)
	for i := 0; i < n; i++ {
		em[i] = make([]int, 0)
	}
	for _, e := range edges {
		em[e[0]] = append(em[e[0]], e[1])
		em[e[1]] = append(em[e[1]], e[0])
	}

	var dfs func(i, p int) []int
	dfs = func(i, p int) []int {
		cnt := make([]int, 26)
		for _, j := range em[i] {
			if j == p {
				continue
			}
			cnt = merge(cnt, dfs(j, i))
		}
		cnt[labels[i]-'a']++
		ans[i] = cnt[labels[i]-'a']
		return cnt
	}

	dfs(0, -1)
	return
}

func merge(a, b []int) []int {
	for i := range a {
		a[i] += b[i]
	}
	return a
}

// @lc code=end
