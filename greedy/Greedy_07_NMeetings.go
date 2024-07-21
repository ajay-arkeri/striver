package greedy

import (
	"sort"
)

// n := 6
// start := []int{1, 3, 0, 5, 8, 5}
// end := []int{2, 4, 5, 7, 9, 9}

type meetings struct {
	start    int
	end      int
	position int
}

//TC - O(Nlogn) + O(N)

func NMeetingsinRoom(start []int, end []int, n int) ([]int,int) {

	meets := make([]meetings, n)

	for i := 0; i < n; i++ {
		meets[i] = meetings{start: start[i], end: end[i], position: i + 1}
	}

	sort.Slice(meets,func(i, j int) bool {
		return meets[i].end < meets[j].end
	})
     
	answer:=[]int{}

	answer = append(answer, meets[0].position)
	lastEndingTime:= meets[0].end

	for i:=1;i<len(meets);i++{
		if meets[i].start>lastEndingTime{
			answer = append(answer, meets[i].position)
			lastEndingTime = meets[i].end
		}
	}

	return answer,len(answer)

}


// leetcode 646 
func FindLongestChain(pairs [][]int) int {
    sort.Slice(pairs,func(i, j int) bool {
		return pairs[i][1]<pairs[j][1]
	})

	previousEnding:=pairs[0][1]
	count:=1
    
	for i:=1;i<len(pairs);i++{
		if pairs[i][0]>previousEnding{
			previousEnding = pairs[i][1]
			count++
		}
	}
    
	return count
}