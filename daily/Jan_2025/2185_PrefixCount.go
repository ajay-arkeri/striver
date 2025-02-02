package daily

// words := []string{"pay", "attention", "practice", "attend"}
func PrefixCount(words []string, pref string) int {
	ans := 0

	for _, word := range words {
		if len(word) >= len(pref) && word[:len(pref)] == pref {
			ans++
		}
	}

	return ans
}
