package slidingwindow

//TC - O(N*N)
//SC - O(256) / O(N)
func LengthOfLongestSubstring_Brute(s string) int {
	n := len(s)

	maxLen := 0

	for i := 0; i < n; i++ {
		mp := make(map[byte]bool)
		for j := i; j < n; j++ {
			if _, ok := mp[s[j]]; ok {
				break
			}
			mp[s[j]] = true
			maxLen = max(maxLen, j-i+1)
		}
	}

	return maxLen
}

//TC - O(N)
//SC - O(N)

func LengthOfLongestSubstring_Optimal(s string) int {
	n := len(s)

	maxLen := 0

	l,r:=0,0
    mp:=make(map[byte]int)

	for r<n{
        if index,ok:=mp[s[r]];ok{
			if index>=l{
				l = index + 1 
			}
		}
		mp[s[r]] = r 
		maxLen = max(maxLen,r-l+1)
		r++
	
	}

	return maxLen
}


