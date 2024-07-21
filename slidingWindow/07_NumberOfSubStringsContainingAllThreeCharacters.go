package slidingwindow

//TC - O(N*N)
//SC - O(1)
func NumberOfSubstrings_Brute(s string) int {
	n := len(s)

	count := 0

	for i := 0; i < n; i++ {
		mp := map[byte]int{'a': 0, 'b': 0, 'c': 0}
		for j := i; j < n; j++ {
			mp[s[j]] = 1
			if mp['a']+mp['b']+mp['c'] == 3 {
				count++
			}
		}
	}

	return count
}

// Once we find a valid substring, just add remaning letters to the count
//TC - O(N*N) at worst case like "aaaaaaa"
//SC - O(1)
func NumberOfSubstrings_Better(s string) int {
	n := len(s)

	count := 0

	for i := 0; i < n; i++ {
		mp := map[byte]int{'a': 0, 'b': 0, 'c': 0}
		for j := i; j < n; j++ {
			mp[s[j]] = 1
			if mp['a']+mp['b']+mp['c'] == 3 {
				count += n - j
				break
			}
		}
	}

	return count
}

// With every character there is a substring that ends
// TC -O(N)
// SC - O(1)
func NumberOfSubstrings_Optimal(s string) int {
	n := len(s)

	count := 0
	lastSeen := []int{-1, -1, -1}

	for i := 0; i < n; i++ {
		lastSeen[s[i]-'a'] = i

		if lastSeen[0] != -1 && lastSeen[1] != -1 && lastSeen[2] != -1 {
			count += 1 + min(lastSeen[0], lastSeen[1], lastSeen[2])
		}
	}

	return count
}

// TC - O(N)
// SC - O(1)
func NumberOfSubstrings_Optimal2(s string) int {
	n := len(s)

	count := 0
	lastSeen := []int{-1, -1, -1}

	for i := 0; i < n; i++ {
		lastSeen[s[i]-'a'] = i

		count += 1 + min(lastSeen[0], lastSeen[1], lastSeen[2])

	}

	return count
}