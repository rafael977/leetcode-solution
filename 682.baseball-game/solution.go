package baseballgame

import "strconv"

/*
 * @lc app=leetcode id=682 lang=golang
 *
 * [682] Baseball Game
 */

// @lc code=start
func calPoints(ops []string) int {
	scores := []int{}
	for _, o := range ops {
		switch o {
		case "C":
			scores = scores[:len(scores)-1]
		case "D":
			s := 2 * scores[len(scores)-1]
			scores = append(scores, s)
		case "+":
			s := scores[len(scores)-2] + scores[len(scores)-1]
			scores = append(scores, s)
		default:
			s, _ := strconv.Atoi(o)
			scores = append(scores, s)
		}
	}

	r := 0
	for _, s := range scores {
		r += s
	}
	return r
}

// @lc code=end
