package greedy

import "sort"

//TC = O(n) + O(nlogn)
//SC = O(1)

func EraseOverlapIntervals(intervals [][]int) int {
    
	n:=len(intervals)

	sort.Slice(intervals,func(i, j int) bool {
		return intervals[i][1]<intervals[j][1]
	})

	count:=1
	lastEndingTime:= intervals[0][1]

	for i:=1;i<n;i++{
       if intervals[i][0]>=lastEndingTime{
		  lastEndingTime = intervals[i][1]
		  count++
	   }
	}

	return n - count	
}