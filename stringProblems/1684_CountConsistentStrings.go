package stringProblems

func CountConsistentStrings(allowed string, words []string) int {
	result := 0

	distinctMap := make(map[byte]bool)

	for i := range allowed {
		distinctMap[allowed[i]] = true
	}

	for _, word := range words {
		flag := true
		for i := range word {
			if !distinctMap[word[i]] {
				flag = false
				break
			}
		}

		if flag {
			result += 1
		}

	}

	return result
}
