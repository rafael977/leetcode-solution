package jumpgameiv

/*
 * @lc app=leetcode id=1345 lang=golang
 *
 * [1345] Jump Game IV
 */

// @lc code=start
func minJumps(arr []int) (ans int) {
	n := len(arr)

	visited := make([]bool, n)
	visited[0] = true
	vtoi := make(map[int][]int)
	for i, v := range arr {
		if _, ok := vtoi[v]; !ok {
			vtoi[v] = make([]int, 0)
		}
		vtoi[v] = append(vtoi[v], i)
	}

	pos := []int{0}
	for len(pos) > 0 {
		nextPos := make([]int, 0)
		for _, p := range pos {
			if p == n-1 {
				return
			}

			if p-1 >= 0 && !visited[p-1] {
				nextPos = append(nextPos, p-1)
				visited[p-1] = true
			}
			if p+1 < n && !visited[p+1] {
				nextPos = append(nextPos, p+1)
				visited[p+1] = true
			}
			for _, j := range vtoi[arr[p]] {
				if !visited[j] {
					nextPos = append(nextPos, j)
					visited[j] = true
				}
			}

		}
		ans++
		pos = nextPos
	}
	return 0
}

// @lc code=end
