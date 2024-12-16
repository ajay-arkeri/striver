package daily

import "container/heap"

type clazz struct {
	change float64
	pass   float64
	total  float64
}

type maxHeap []clazz

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].change > h[j].change }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(clazz))
}

func MaxAverageRatio(classes [][]int, extraStudents int) float64 {
	var result float64
	h := make(maxHeap, len(classes))

	for i := range classes {
		h[i] = clazz{
			pass:  float64(classes[i][0]),
			total: float64(classes[i][1]),
		}
		h[i].change = ((h[i].pass + 1) / (h[i].total + 1)) - (h[i].pass / h[i].total)
	}

	heap.Init(&h)

	for i := 0; i < extraStudents; i++ {
		clazz := heap.Pop(&h).(clazz)
		clazz.pass++
		clazz.total++
		clazz.change = ((clazz.pass + 1) / (clazz.total + 1)) - (clazz.pass / clazz.total)

		heap.Push(&h, clazz)
	}

	//remove elements from heap and calculate total average
	for i := 0; i < len(h); i++ {
		result += h[i].pass / h[i].total
	}

	return result / float64(len(h))
}
