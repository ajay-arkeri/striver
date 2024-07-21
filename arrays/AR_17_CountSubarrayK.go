package arrays

func SubarraySum_Brute(arr []int, k int) int {
	n := len(arr)
	count := 0

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += arr[j]

			if sum == k {
				count++
			}
		}
	}

	return count
}

// TC - O(N) --> if worst case collision then add O(N)
// SC - O(N)
func SubarraySum_Optimal(arr []int, k int) int {
	n := len(arr)
	count := 0

	prefixSumCount := make(map[int]int)

	prefixSumCount[0] = 1

	sum := 0

	for i := 0; i < n; i++ {
		sum += arr[i]

		if val, ok := prefixSumCount[sum-k]; ok {
			count += val
		}

		prefixSumCount[sum]++
	}

	return count
}