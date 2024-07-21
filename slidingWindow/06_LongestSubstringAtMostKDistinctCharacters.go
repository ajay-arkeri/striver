package slidingwindow

//leetcode - 340,159 (locked)

//TC - O(N*N)
//SC - O(26)
func LengthOfLongestSubstringKDistinct_Brute(s string, k int) (ans int) {
	n := len(s)
	maxCount := 0

	for i := 0; i < n; i++ {
		mp := make(map[byte]int)

		for j := i; j < n; j++ {
			mp[s[j]]++

			if len(mp) > k {
				break
			}

			maxCount = max(maxCount, j-i+1)
		}
	}

	return maxCount
}

//TC - O(N+N)
//SC - O(26)
func LengthOfLongestSubstringKDistinct_Better(s string, k int) (ans int) {
	n := len(s)
	maxCount := 0

	mp:=make(map[byte]int)
    l,r:=0,0

	for r<n{
		mp[s[r]]++

		//shrink
        for len(mp)>k{
		   mp[s[l]]--

		   if mp[s[l]]==0{
			delete(mp,s[l])
		   }

		   l++
		}

		maxCount = max(maxCount,r-l+1)
		r++
	}

	return maxCount
}


//TC - O(N)
//SC - O(26)
func LengthOfLongestSubstringKDistinct_Optimal(s string, k int) (ans int) {
	n := len(s)
	maxCount := 0

	mp:=make(map[byte]int)
    l,r:=0,0

	for r<n{
		mp[s[r]]++
        
		if len(mp)>k{
			mp[s[l]]--
			if mp[s[l]]==0{
				delete(mp,s[l])
			}
			l++
		}
		 

		if len(mp)<=k{
			maxCount= max(maxCount,r-l+1)
		}
		
		r++
	}

	return maxCount
}
