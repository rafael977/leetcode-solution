package longestconsecutivesequence

/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
func longestConsecutive(nums []int) (ans int) {
	nmap := make(map[int]bool)
	for _, v := range nums {
		nmap[v] = true
	}

	for k := range nmap {
		if nmap[k-1] {
			continue
		}

		streak := 1
		n := k + 1
		for nmap[n] {
			streak++
			n++
		}

		if streak > ans {
			ans = streak
		}
	}
	return
}

// @lc code=end
