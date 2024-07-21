package greedy

func InsertInterval(intervals [][]int, newInterval []int) [][]int {

	answer := [][]int{}

	i := 0
	//Left
	for intervals[i][1] < newInterval[0] {
		answer = append(answer, intervals[i])
		i++
	}

	//middle
	for intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}

	answer = append(answer, newInterval)

	//Right
	for i < len(intervals) {
		answer = append(answer, intervals[i])
		i++
	}

	return answer
}