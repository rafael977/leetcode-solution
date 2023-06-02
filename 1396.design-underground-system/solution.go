package designundergroundsystem

/*
 * @lc app=leetcode id=1396 lang=golang
 *
 * [1396] Design Underground System
 */

// @lc code=start
type travelling struct {
	in   string
	time int
}

type aggregate struct {
	sum, count int
}

type UndergroundSystem struct {
	t map[int]travelling
	a map[string]map[string]aggregate
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		t: make(map[int]travelling),
		a: make(map[string]map[string]aggregate),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.t[id] = travelling{in: stationName, time: t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	tvl := this.t[id]
	delete(this.t, id)
	if _, ok := this.a[tvl.in]; !ok {
		this.a[tvl.in] = make(map[string]aggregate)
		this.a[tvl.in][stationName] = aggregate{sum: 0, count: 0}
	}
	agg := this.a[tvl.in][stationName]
	agg.sum += t - tvl.time
	agg.count++
	this.a[tvl.in][stationName] = agg
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	return float64(this.a[startStation][endStation].sum) / float64(this.a[startStation][endStation].count)
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
// @lc code=end
