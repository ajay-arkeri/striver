package daily

import (
	"container/heap"
)

type charPair struct {
	char  byte
	count int
}

type maxHeapChar []charPair

func (h maxHeapChar) Len() int           { return len(h) }
func (h maxHeapChar) Less(i, j int) bool { return h[i].char > h[j].char }
func (h maxHeapChar) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeapChar) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func (h *maxHeapChar) Push(x any) {
	*h = append(*h, x.(charPair))
}

func RepeatLimitedString_Heap(s string, repeatLimit int) string {
	result := []byte{}
	freq := map[byte]int{}

	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	h := maxHeapChar{}

	for char, count := range freq {
		h = append(h, charPair{char: char, count: count})
	}

	heap.Init(&h)

	for len(h) > 0 {
		currentPair := heap.Pop(&h).(charPair)

		k := min(currentPair.count, repeatLimit)

		for i := 0; i < k; i++ {
			result = append(result, currentPair.char)
		}
		currentPair.count -= k

		if currentPair.count > 0 {
			if len(h) == 0 {
				break
			}
			//add a separator
			nextPair := heap.Pop(&h).(charPair)
			result = append(result, nextPair.char)
			nextPair.count--
			if nextPair.count > 0 {
				heap.Push(&h, nextPair)
			}
			heap.Push(&h, currentPair)
		}
	}

	return string(result)
}

func RepeatLimitedString_Array(s string, repeatLimit int) string {

	freq := make([]int, 26)

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	index := 25
	result := []byte{}

	for index >= 0 {
		if freq[index] == 0 {
			index--
			continue
		}

		k := min(repeatLimit, freq[index])

		for i := 0; i < k; i++ {
			result = append(result, byte(index+'a'))
		}

		freq[index] -= k
		lettersLeft := freq[index]
		if lettersLeft > 0 {
			prevIndex := index - 1

			for prevIndex >= 0 && freq[prevIndex] == 0 {
				prevIndex--
			}

			if prevIndex == -1 {
				break
			}
			result = append(result, byte(prevIndex+'a'))
			freq[prevIndex]--
		}

	}

	return string(result)
}
