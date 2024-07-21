package slidingwindow

import (
	"math"
)

//TC - O(N*(N+M)) ~ O(N^2)
//SC  - O(M)
func MinWindow_Brute(s string, t string) string {
	sLen := len(s)
	tLen := len(t)
	minLen := math.MaxInt
    startIndex:=-1

	for i := 0; i < sLen; i++ {
		tMap := make(map[byte]int)
		for i := range t {
			tMap[t[i]]++
		}
        count:=0 

		for j := i; j < sLen; j++ {
            if tMap[s[j]]>0{
				count++
			}
			tMap[s[j]]--

			if count==tLen{
                if j-i+1<minLen{
					startIndex = i
					minLen = j-i+1
				}
				break	
			}
		}
	}

	if startIndex==-1{
		return ""
	}else{
		return s[startIndex:startIndex+minLen]
	}
}

//TC - O(M+2N)
//SC - O(M)
func MinWindow_Optimal(s string, t string) string {
	sLen := len(s)
	tLen := len(t)
	minLen := math.MaxInt
    startIndex:=-1
	count:=0 

	tMap := make(map[byte]int)
	for i := range t {
		tMap[t[i]]++
	}

	l,r:=0,0

	for r<sLen {
	    if tMap[s[r]]>0{
			count++
		}

		tMap[s[r]]--

		for count==tLen{
           if r-l+1<minLen{
			  minLen = r-l+1
			  startIndex = l
		   }
           
		   tMap[s[l]]++

		   if tMap[s[l]]>0{
			count--
		   }
           
		   l++
		}

		r++
	}

	if startIndex==-1{
		return ""
	}else{
		return s[startIndex:startIndex+minLen]
	}
}