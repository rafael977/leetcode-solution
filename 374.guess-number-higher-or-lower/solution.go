package guessnumberhigherorlower

/*
 * @lc app=leetcode id=374 lang=golang
 *
 * [374] Guess Number Higher or Lower
 */

var guess func(num int) int

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	lo, hi := 1, n
	m := lo + (hi-lo)/2
	r := guess(m)
	for r != 0 {
		if r == -1 {
			hi = m - 1
		} else if r == 1 {
			lo = m + 1
		}
		m = lo + (hi-lo)/2
		r = guess(m)
	}

	return m
}

// @lc code=end
