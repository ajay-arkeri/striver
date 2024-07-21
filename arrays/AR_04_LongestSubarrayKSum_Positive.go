package arrays

import (
	"math"
)

func LargestSubarrayWith_K_Sum_Positives_Brute(arr []int, k int) int {
	n := len(arr)

	largest := math.MinInt

	for i:=0;i<n;i++{
        sum:=0
		for j:=i;j<n;j++{
             sum+=arr[j]
            
			 if sum==k{
				largest = max(largest, j-i+1)
				break
			 }else if sum>k{
				break
			 }
		}
	}

	return largest

}


//prefixSum using map
//This the Optimal solution for negatives 
//TC -- O(N)  SC-- O(N)
func LargestSubarrayWith_K_Sum_Positives_Better(arr []int, k int) int {
	n := len(arr)
      
    sum,maxLen:=0,0

	mp:= map[int]int{}

	for i:=0;i<n;i++{
		sum+=arr[i]

		if sum==k{
			maxLen = max(maxLen,i+1)
		}

		if len(mp)!=0{
			if val,exits:= mp[sum-k];exits{
				maxLen = max(maxLen,i-val)
			}
		}

		if _,exists:= mp[sum-k];!exists{
             mp[sum]=i
		}
	
	}

	return maxLen

}

//2 pointer --> increase j pointer and reduce i pointer.
//TC --> O(2N) --> not O(N^2) because inside loop ran for max N times
func LargestSubarrayWith_K_Sum_Positives_Optimal(arr []int, k int) int {
	n := len(arr)
      
    sum,maxLen:=0,0
    
	i,j:=0,0
   
	for j<n{
       sum+=arr[j]
	   
	   for i<=j && sum>k{
		  sum-=arr[i]
		  i++
	   } 

	   if sum==k{
		 maxLen = max(maxLen,j-i+1)
	   }
	   j++
	}
    
	return maxLen
}
