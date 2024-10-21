package checkifitisastraightline

/*
 * @lc app=leetcode id=1232 lang=golang
 *
 * [1232] Check If It Is a Straight Line
 */

// @lc code=start
func checkStraightLine(coordinates [][]int) bool {
	p1, p2 := coordinates[0], coordinates[1]
	a, b, k := p1[1]-p2[1], p1[0]*p2[1]-p1[1]*p2[0], p1[0]-p2[0]

	for _, v := range coordinates[2:] {
		l, r := a*v[0]+b, v[1]*k
		if l != r {
			return false
		}
	}

	return true
}

// @lc code=end
