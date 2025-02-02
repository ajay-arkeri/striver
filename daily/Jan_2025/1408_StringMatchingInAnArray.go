package daily

import "strings"

// words := []string{"leetcoder", "leetcode", "od", "hamlet", "am"}

func StringMatching(words []string) []string {
	n := len(words)

	mp := map[string]bool{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			if strings.Contains(words[j], words[i]) {
				if !mp[words[i]] {
					mp[words[i]] = true
				}
			}
		}
	}

	ans := []string{}

	for word := range mp {
		ans = append(ans, word)
	}

	return ans
}

func StringMatching_Optimal(words []string) []string {
	n := len(words)

	ans := []string{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			if strings.Contains(words[j], words[i]) {
				ans = append(ans, words[i])
				break
			}
		}
	}

	return ans
}
