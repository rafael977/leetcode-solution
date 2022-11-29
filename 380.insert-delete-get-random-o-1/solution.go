package insertdeletegetrandomo1

import "math/rand"

/*
 * @lc app=leetcode id=380 lang=golang
 *
 * [380] Insert Delete GetRandom O(1)
 */

// @lc code=start
type RandomizedSet struct {
    valToIdx map[int]int
	vals []int
}


func Constructor() RandomizedSet {
    return RandomizedSet{
		valToIdx: make(map[int]int),
		vals: make([]int, 0),
	}
}


func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valToIdx[val]; ok {
		return false
	}
	this.valToIdx[val] = len(this.vals)
	this.vals = append(this.vals, val)
	return true
}


func (this *RandomizedSet) Remove(val int) bool {
    if _, ok := this.valToIdx[val]; !ok {
		return false
	}
	last := len(this.vals)-1
	this.vals[this.valToIdx[val]] = this.vals[last]
	this.valToIdx[this.vals[last]] = this.valToIdx[val]
	delete(this.valToIdx, val)
	this.vals = this.vals[:last]
	return true
}


func (this *RandomizedSet) GetRandom() int {
    idx := rand.Intn(len(this.vals))
	return this.vals[idx]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end

