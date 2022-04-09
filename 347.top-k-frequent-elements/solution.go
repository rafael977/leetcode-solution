package topkfrequentelements

import "sort"

/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
type KV struct {
	Key   int
	Value int
}

func topKFrequent(nums []int, k int) (ans []int) {
	f := make(map[int]int)
	for _, v := range nums {
		f[v]++
	}

	kvs := make([]KV, len(f))
	for k, v := range f {
		kvs = append(kvs, KV{Key: k, Value: v})
	}

	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].Value > kvs[j].Value
	})

	for i := 0; i < k; i++ {
		ans = append(ans, kvs[i].Key)
	}
	return
}

// @lc code=end
