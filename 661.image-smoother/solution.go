package imagesmoother

/*
 * @lc app=leetcode id=661 lang=golang
 *
 * [661] Image Smoother
 */

// @lc code=start
var dir = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 0}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])

	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for i := range img {
		for j := range img[i] {
			sum := 0
			c := 0
			for _, d := range dir {
				ii, jj := i+d[0], j+d[1]
				if ii >= 0 && ii < m && jj >= 0 && jj < n {
					sum += img[ii][jj]
					c++
				}
			}
			ans[i][j] = sum / c
		}
	}
	return ans
}

// @lc code=end
