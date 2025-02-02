package daily

import "strings"

func MaxScore_Brute(s string) int {
	ans := 0

	n := len(s)

	for i := 1; i < n; i++ {
		score := 0

		leftSubstring := s[:i]
		rightSubstring := s[i:]

		for _, ch := range leftSubstring {
			if ch == '0' {
				score++
			}
		}

		for _, ch := range rightSubstring {
			if ch == '1' {
				score++
			}
		}

		ans = max(ans, score)
	}

	return ans
}

func MaxScore_Optimal(s string) int {
	ans := 0

	ones := strings.Count(s, "1")
	zeros := 0

	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			zeros++
		} else {
			ones--
		}

		ans = max(ans, zeros+ones)
	}

	return ans

}
