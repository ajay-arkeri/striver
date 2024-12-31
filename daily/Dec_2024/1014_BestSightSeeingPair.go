package daily

func MaxScoreSightseeingPair_Brute(values []int) int {
	ans := 0
	n := len(values)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ans = max(values[j]+values[i]+i-j, ans)
		}
	}
	return ans
}

func MaxScoreSightseeingPair_Optimal(values []int) int {
	n := len(values)
	ans := 0
	curMax := values[0] - 1

	for i := 1; i < n; i++ {
		ans = max(curMax+values[i], ans)
		curMax = max(curMax-1, values[i]-1)
	}

	return ans
}
