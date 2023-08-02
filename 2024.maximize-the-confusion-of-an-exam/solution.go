package maximizetheconfusionofanexam

/*
 * @lc app=leetcode id=2024 lang=golang
 *
 * [2024] Maximize the Confusion of an Exam
 */

// @lc code=start
func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(maxConsecutiveChar(answerKey, k, 'T'), maxConsecutiveChar(answerKey, k, 'F'))
}

func maxConsecutiveChar(answerKey string, k int, c byte) (ans int) {
	l := 0
	sum := 0
	for r := range answerKey {
		if answerKey[r] != c {
			sum++
		}
		for sum > k {
			if answerKey[l] != c {
				sum--
			}
			l++
		}
		ans = max(ans, r-l+1)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
