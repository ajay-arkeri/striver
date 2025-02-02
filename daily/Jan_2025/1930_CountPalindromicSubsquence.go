package daily

func CountPalindromicSubsequence_Brute(s string) int {
	n := len(s)

	type pair struct {
		outer  byte
		middle byte
	}

	mp := make(map[pair]bool)

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if s[i] == s[k] {
					mp[pair{s[i], s[j]}] = true
				}
			}
		}
	}

	return len(mp)
}

func CountPalindromicSubsequence_Optimal(s string) int {
	n := len(s)

	type pair struct {
		outer  byte
		middle byte
	}

	res := make(map[pair]bool)

	right := make(map[byte]int)
	left := make(map[byte]bool)

	for i := 0; i < n; i++ {
		right[s[i]]++
	}

	for i := 0; i < n; i++ {

		right[s[i]]--
		if right[s[i]] == 0 {
			delete(right, s[i])
		}
		//compare
		for key := range left {
			if _, ok := right[key]; ok {
				res[pair{outer: key, middle: s[i]}] = true
			}
		}

		//move the middle
		left[s[i]] = true

	}

	return len(res)
}
