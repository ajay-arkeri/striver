package daily

// words := []string{"a", "aba", "ababa", "aa"}

func CountPrefixSuffixPairs(words []string) int {
	n := len(words)
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if i == j || i > j {
				continue
			}

			if len(words[i]) <= len(words[j]) && words[j][:len(words[i])] == words[i] && words[j][len(words[j])-len(words[i]):] == words[i] {
				ans++
			}

		}
	}
	return ans
}
