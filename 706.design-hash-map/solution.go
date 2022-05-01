package designhashmap

/*
 * @lc app=leetcode id=706 lang=golang
 *
 * [706] Design HashMap
 */

// @lc code=start
const mod int = 997

type entry struct {
	key, value int
}

type MyHashMap struct {
	m [][]entry
}

func Constructor() MyHashMap {
	return MyHashMap{m: make([][]entry, mod)}
}

func (this *MyHashMap) Put(key int, value int) {
	i := key % mod
	if this.m[i] == nil {
		this.m[i] = make([]entry, 0)
	}
	for j := range this.m[i] {
		if this.m[i][j].key == key {
			this.m[i][j].value = value
			return
		}
	}
	this.m[i] = append(this.m[i], entry{key: key, value: value})
}

func (this *MyHashMap) Get(key int) int {
	i := key % mod
	if this.m[i] == nil {
		return -1
	}
	for _, v := range this.m[i] {
		if v.key == key {
			return v.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	i := key % mod
	if this.m[i] == nil {
		return
	}
	for j, v := range this.m[i] {
		if v.key == key {
			this.m[i] = append(this.m[i][:j], this.m[i][j+1:]...)
			return
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
// @lc code=end
