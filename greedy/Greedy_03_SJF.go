package greedy

import "sort"

//TC - O(N)
//SC - O(1)

func AverageWaitTime_SJF(jobs []int) int {
	n := len(jobs)

	waitTime := 0
	startTime := 0

	sort.Ints(jobs)

	for i := 0; i < n; i++ {
		waitTime += startTime
		startTime += jobs[i]
	}

	return waitTime / n
}