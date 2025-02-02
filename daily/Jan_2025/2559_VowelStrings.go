package daily

import "strings"

// words := []string{"aba", "bcb", "ece", "aa", "e"}
// queries := [][]int{{0, 2}, {1, 4}, {1, 1}}

func VowelStrings(words []string, queries [][]int) []int {
	prefixArray := make([]int, len(words)+1)
	ans := make([]int, len(queries))

	vowels := "aeiou"
	sum := 0
	for i, word := range words {
		if strings.ContainsRune(vowels, rune(word[0])) && strings.ContainsRune(vowels, rune(word[len(word)-1])) {
			prefixArray[i+1] = sum + 1
			sum++
		} else {
			prefixArray[i+1] = sum
		}
	}

	for i, query := range queries {
		l, r := query[0], query[1]
		ans[i] = prefixArray[r+1] - prefixArray[l]
	}

	return ans
}
