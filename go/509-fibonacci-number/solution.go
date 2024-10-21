package fibonaccinumber

/*
 * @lc app=leetcode id=509 lang=golang
 *
 * [509] Fibonacci Number
 */

// @lc code=start
func fib(n int) int {
	nmap := make(map[int]int)
	var rec func(n int) int
	rec = func(n int) int {
		if n == 0 || n == 1 {
			return n
		}

		var n1, n2 int
		var ok bool
		if n1, ok = nmap[n-1]; !ok {
			n1 = rec(n - 1)
		}
		if n2, ok = nmap[n-2]; !ok {
			n2 = rec(n - 2)
		}

		nmap[n] = n1 + n2
		return nmap[n]
	}
	return rec(n)
}

// @lc code=end
