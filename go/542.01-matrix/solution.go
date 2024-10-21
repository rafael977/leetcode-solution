package zeroonematrix

/*
 * @lc app=leetcode id=542 lang=golang
 *
 * [542] 01 Matrix
 */

// @lc code=start
func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	q := make([][]int, 0)
	for i, r := range mat {
		for j, v := range r {
			if v == 0 {
				q = append(q, []int{i, j})
				visited[i][j] = true
			}
		}
	}

	for len(q) > 0 {
		i, j := q[0][0], q[0][1]
		q = q[1:]
		if i-1 >= 0 && !visited[i-1][j] {
			mat[i-1][j] = mat[i][j] + 1
			visited[i-1][j] = true
			q = append(q, []int{i - 1, j})
		}
		if j-1 >= 0 && !visited[i][j-1] {
			mat[i][j-1] = mat[i][j] + 1
			visited[i][j-1] = true
			q = append(q, []int{i, j - 1})
		}
		if i+1 < m && !visited[i+1][j] {
			mat[i+1][j] = mat[i][j] + 1
			visited[i+1][j] = true
			q = append(q, []int{i + 1, j})
		}
		if j+1 < n && !visited[i][j+1] {
			mat[i][j+1] = mat[i][j] + 1
			visited[i][j+1] = true
			q = append(q, []int{i, j + 1})
		}
	}

	return mat
}

// @lc code=end
