package miscellaneous

import "sort"

//Leetcode Problem number - 3014

//not distinct
func MinimumPushes_Brute1(word string) int {
    count := 0

	for i := range word {
		count += (i / 8) + 1
	}

	return count
}


func MinimumPushes_Brute2(word string) int {
    n:=len(word)

    if n<8{
        return n
    }

    if n<16{
        return 8 + 2*(n-8)
    }

    if n<24{
        return 8 + 16 + 3*(n-16)
    }

    return 8 + 16 + 24 + 4 *(n-24)
}

//Distinct
func MinimumPushes(word string) int {
   	ans := 0

	occurenceArray := make([]int,27)
	     
	for _, letter := range word {
		occurenceArray[int(letter)-96]++
	}
      
	 sort.Slice(occurenceArray,func(i, j int) bool {
		return occurenceArray[i]>occurenceArray[j]
	 })

	 for i,val:=range occurenceArray{
		 count:=(i/8)+1
         
		 if val>0{
           ans+= count*val
		 }

	 }

	 return ans
}