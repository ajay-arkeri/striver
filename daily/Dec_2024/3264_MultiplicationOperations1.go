package daily

import (
	"container/heap"
	"fmt"
)

type pair struct {
	val   int
	index int
}

type minHeap []pair

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Less(i, j int) bool {
	if h[i].val == h[j].val {
		return h[i].index < h[j].index
	}
	return h[i].val < h[j].val
}
func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(pair))
}

func GetFinalState(nums []int, k int, multiplier int) []int {
	h := &minHeap{}
	for i, val := range nums {
		*h = append(*h, pair{val: val, index: i})
	}

	heap.Init(h)

	for i := 0; i < k; i++ {
		pair := heap.Pop(h).(pair)
		fmt.Println(nums)
		pair.val = pair.val * multiplier
		nums[pair.index] = pair.val
		heap.Push(h, pair)
	}

	return nums
}
