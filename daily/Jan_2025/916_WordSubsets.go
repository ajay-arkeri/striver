package daily

import "fmt"

// words1 := []string{"amazon", "apple", "facebook", "google", "leetcode"}
// words2 := []string{"e", "a"}

func getFrequencyCount(word string) map[byte]int {
	mp := make(map[byte]int)

	for i := 0; i < len(word); i++ {
		mp[word[i]]++
	}

	return mp
}

func WordSubsets(words1 []string, words2 []string) []string {
	ans := []string{}

	targetMp := make(map[byte]int)

	for _, word := range words2 {
		subFreqMp := getFrequencyCount(word)

		fmt.Println(subFreqMp)

		for key, val := range subFreqMp {
			if targetMp[key] < val {
				targetMp[key] = val
			}
		}
	}

	for _, word := range words1 {
		freqMp := getFrequencyCount(word)

		isUniversal := true

		for key, val := range targetMp {
			if freqMp[key] < val {
				isUniversal = false
				break
			}
		}

		if isUniversal {
			ans = append(ans, word)
		}
	}

	return ans
}
