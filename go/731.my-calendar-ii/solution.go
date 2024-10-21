package mycalendarii

/*
 * @lc app=leetcode id=731 lang=golang
 *
 * [731] My Calendar II
 */

// @lc code=start
type MyCalendarTwo struct {
	booked, overlap [][2]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		booked:  make([][2]int, 0),
		overlap: make([][2]int, 0),
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, p := range this.overlap {
		if p[0] < end && start < p[1] {
			return false
		}
	}

	for _, b := range this.booked {
		if b[0] < end && start < b[1] {
			this.overlap = append(this.overlap, [2]int{max(b[0], start), min(b[1], end)})
		}
	}
	this.booked = append(this.booked, [2]int{start, end})
	return true
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
// @lc code=end
