package daily

// a := []int{1, 3, 2, 4}
// b := []int{3, 1, 4, 2}
func FindThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	freq := make([]int, n+1)

	ans := make([]int, n)

	count := 0
	for i := 0; i < n; i++ {
		freq[A[i]]++

		if freq[A[i]] == 2 {
			count++
		}

		freq[B[i]]++

		if freq[B[i]] == 2 {
			count++
		}

		ans[i] = count
	}

	return ans
}
