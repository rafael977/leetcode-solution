package timebasedkeyvaluestore

/*
 * @lc app=leetcode id=981 lang=golang
 *
 * [981] Time Based Key-Value Store
 */

// @lc code=start
type TimeMap struct {
	m map[string][]struct {
		ts int
		v  string
	}
}

func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][]struct {
			ts int
			v  string
		}),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if this.m[key] == nil {
		this.m[key] = make([]struct {
			ts int
			v  string
		}, 0)
	}
	this.m[key] = append(this.m[key], struct {
		ts int
		v  string
	}{ts: timestamp, v: value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	list := this.m[key]
	if list == nil {
		return ""
	}
	for i := len(list) - 1; i >= 0; i-- {
		if timestamp >= list[i].ts {
			return list[i].v
		}
	}
	return ""
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
// @lc code=end
