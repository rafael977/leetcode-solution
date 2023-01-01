package singlethreadedcpu

import (
	"container/heap"
	"sort"
)

/*
 * @lc app=leetcode id=1834 lang=golang
 *
 * [1834] Single-Threaded CPU
 */

// @lc code=start
type task struct {
	i, st, pt int
}

type taskHeap struct {
	tasks []task
}

func (h taskHeap) Len() int {
	return len(h.tasks)
}

func (h taskHeap) Less(i, j int) bool {
	if h.tasks[i].pt == h.tasks[j].pt {
		return h.tasks[i].i < h.tasks[j].i
	}
	return h.tasks[i].pt < h.tasks[j].pt
}

func (h *taskHeap) Swap(i, j int) {
	h.tasks[i], h.tasks[j] = h.tasks[j], h.tasks[i]
}

func (h *taskHeap) Push(x interface{}) {
	h.tasks = append(h.tasks, x.(task))
}

func (h *taskHeap) Pop() interface{} {
	old := h.tasks
	x := old[h.Len()-1]
	h.tasks = h.tasks[:h.Len()-1]
	return x
}

func getOrder(tasks [][]int) (ans []int) {
	tt := make([]task, len(tasks))
	for i, v := range tasks {
		tt[i] = task{i: i, st: v[0], pt: v[1]}
	}
	sort.Slice(tt, func(i, j int) bool {
		if tt[i].st == tt[j].st {
			return tt[i].pt < tt[j].pt
		}
		return tt[i].st < tt[j].st
	})

	thp := &taskHeap{}
	heap.Init(thp)
	heap.Push(thp, tt[0])
	time := tt[0].st
	j := 1
	for thp.Len() > 0 || j < len(tt) {
		if thp.Len() == 0 {
			time = tt[j].st
			for j < len(tt) && time >= tt[j].st {
				heap.Push(thp, tt[j])
				j++
			}
			continue
		}
		t := heap.Pop(thp).(task)
		ans = append(ans, t.i)
		time += t.pt
		for j < len(tt) && time >= tt[j].st {
			heap.Push(thp, tt[j])
			j++
		}
	}
	return
}

// @lc code=end
