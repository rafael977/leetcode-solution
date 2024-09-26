package mycalendari

import "sort"

/*
 * @lc app=leetcode id=729 lang=golang
 *
 * [729] My Calendar I
 */

// @lc code=start
type MyCalendar struct {
	s [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{s: make([][2]int, 0)}
}

func (this *MyCalendar) Book(start int, end int) bool {
	i := sort.Search(len(this.s), func(i int) bool {
		return end <= this.s[i][0]
	})

	if i <= 0 || this.s[i-1][1] <= start {
		this.s = append(this.s, [2]int{start, end})
		copy(this.s[i+1:], this.s[i:])
		this.s[i] = [2]int{start, end}
		return true
	}
	return false
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
// @lc code=end
