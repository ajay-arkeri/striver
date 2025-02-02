package daily

// s := "annabelle" k=2

func CanConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}

	mp := make(map[byte]int)

	for i := range s {
		mp[s[i]]++
	}

	odd := 0
	for _, val := range mp {
		odd += val % 2
	}

	return odd <= k

}
