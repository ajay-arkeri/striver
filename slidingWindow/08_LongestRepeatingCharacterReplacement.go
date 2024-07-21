package slidingwindow

//"ABAB"
// No of changes required = len of subString - MaxFrequency
func CharacterReplacement_Brute(s string, k int) int {
    n:=len(s)
	maxLen:=0

	for i:=0;i<n;i++{
		maxFreq:=0
		mp:=make([]int,26)

		for j:=i;j<n;j++{
		    mp[s[j]-'A']++
			
			maxFreq = max(maxFreq,mp[s[j]-'A'])

			subStringLen:= j-i+1

			if subStringLen-maxFreq<=k{
				maxLen = max(maxLen,subStringLen)
			}else{
				break	
			}
		}
	}

	return maxLen
}

//While shrinking recalculate the maxFreq
//TC - O(N+N *26)
func CharacterReplacement_Better(s string, k int) int {
    n:=len(s)
	maxLen:=0

	mp:=make([]int,26)

	l,r:=0,0
	maxFreq:=0


	for r<n{
        mp[s[r]-'A']++

		maxFreq = max(maxFreq,mp[s[r]-'A'])
        
		
        
		//shrinking
		for r-l+1 - maxFreq>k{
            mp[s[l]-'A']--
            maxFreq = 0
			//recalculate freq
			for _,val:= range mp{
				maxFreq = max(maxFreq,val)
			}
			l++
		} 
       
		if r-l+1 - maxFreq<=k{
			maxLen  = max(maxLen,r-l+1)
		}
         
		r++
	}

	return maxLen
}


//TC - O(N)
//SC - O(26)
func CharacterReplacement_Optimal(s string, k int) int {
   n:=len(s)
	maxLen:=0

	mp:=make([]int,26)

	l,r:=0,0
	maxFreq:=0


	for r<n{
        mp[s[r]-'A']++

		maxFreq = max(maxFreq,mp[s[r]-'A'])
        
		//shrinking
		if r-l+1 - maxFreq>k{
            mp[s[l]-'A']--
			l++
		} 
       
		if r-l+1 - maxFreq<=k{
			maxLen  = max(maxLen,r-l+1 )
		}

		r++
	}

	return maxLen
}