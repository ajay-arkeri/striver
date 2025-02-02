package daily

func MinimumLength(s string) int {
	freq := make([]int, 26)

	ans := 0

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	for i := 0; i < len(freq); i++ {
		if freq[i] == 0 {
			continue
		} else if freq[i]%2 == 0 {
			ans += 2
		} else {
			ans += 1
		}

	}

	return ans
}
