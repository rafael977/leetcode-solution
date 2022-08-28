package reorderedpowerof2

/*
 * @lc app=leetcode id=869 lang=golang
 *
 * [869] Reordered Power of 2
 */

// @lc code=start
var nmap [][]int

func initNmap() {
	if len(nmap) > 0 {
		return
	}
	for n := 1; n <= 10e9; n = n << 1 {
		nmap = append(nmap, countDigits(n))
	}
}

func reorderedPowerOf2(n int) bool {
	initNmap()
	dm := countDigits(n)
	for _, v := range nmap {
		match := true
		for i := range dm {
			if dm[i] != v[i] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}

	return false
}

func countDigits(n int) []int {
	ans := make([]int, 10)
	for n > 0 {
		ans[n%10]++
		n /= 10
	}
	return ans
}

// @lc code=end
