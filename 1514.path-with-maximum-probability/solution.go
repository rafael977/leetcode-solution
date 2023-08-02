package pathwithmaximumprobability

import "container/heap"

/*
 * @lc app=leetcode id=1514 lang=golang
 *
 * [1514] Path with Maximum Probability
 */

// @lc code=start
type nodeProb struct {
	idx  int
	prob float64
}

type nodeQueue []nodeProb

func (q nodeQueue) Len() int {
	return len(q)
}

func (q nodeQueue) Less(i, j int) bool {
	return q[i].prob > q[j].prob
}

func (q nodeQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *nodeQueue) Push(x interface{}) {
	*q = append(*q, x.(nodeProb))
}

func (q *nodeQueue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	grap := make(map[int]map[int]float64)
	for i, e := range edges {
		if _, ok := grap[e[0]]; !ok {
			grap[e[0]] = make(map[int]float64)
		}
		grap[e[0]][e[1]] = succProb[i]
		if _, ok := grap[e[1]]; !ok {
			grap[e[1]] = make(map[int]float64)
		}
		grap[e[1]][e[0]] = succProb[i]
	}

	q := &nodeQueue{{idx: start, prob: 1}}
	heap.Init(q)

	prob := make([]float64, n)
	prob[start] = 1
	for q.Len() > 0 {
		node := heap.Pop(q).(nodeProb)
		if node.prob < prob[node.idx] {
			continue
		}

		for next, p := range grap[node.idx] {
			nextProb := p * node.prob
			if nextProb > prob[next] {
				prob[next] = nextProb
				heap.Push(q, nodeProb{idx: next, prob: nextProb})
			}
		}
	}

	return prob[end]
}

// @lc code=end
