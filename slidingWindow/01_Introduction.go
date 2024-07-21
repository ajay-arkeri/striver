package slidingwindow

//Patterns

//constant window
//[-1,2,3,4,5,-1] k=4

func MaxSumConsecutive_Pattern(arr []int,k int)int{
	n:=len(arr)

	l:=0;
	r:=k-1
    
	sum:=0

	for i:=l;i<=r;i++{
		sum+= arr[i]
	}

	ans:=0

	ans=max(ans,sum)

	for r<n-1{
		sum-=arr[l]
		l++
		r++
		sum+=arr[r]

		ans = max(ans,sum)
	}

	return ans 
}




 
// Longest subarray / substring where <condition>
//longest subarray with sum <=k
//[2,5,1,7,10]  k=14

func LongestSubarraySum_Pattern_Brute(arr []int, k int)int{
     n:=len(arr)

	 maxAns:=0
     
	 for i:=0;i<n;i++{
        sum:=0
		for j:=i;j<n;j++{
			sum+= arr[j]

			if sum<=k{
				maxAns = max(maxAns,j-i+1)
			}
		}
	 }


	 return maxAns
}

//once sum>k  do i need to continue

func LongestSubarraySum_Pattern_Brute2(arr []int, k int)int{
     n:=len(arr)

	 maxAns:=0
     
	 for i:=0;i<n;i++{
        sum:=0
		for j:=i;j<n;j++{
			sum+= arr[j]
              
            if sum>k{
				break
			}

			if sum<=k{
				maxAns = max(maxAns,j-i+1)
			}
		}
	 }


	 return maxAns
}


//2 pointer expand and shrik the window\
//TC - O(2N)
//SC - O(1)
func LongestSubarraySum_Pattern_Better(arr []int, k int)int{
     n:=len(arr)

	 maxLen:=0
    
	 l,r:=0,0
     sum:=0

	 for r<n{
        sum+= arr[r]

		//shrink
		for sum>k{
			sum-=arr[l]
			l++
		}
          
		if sum<=k{
           maxLen = max(maxLen,r-l+1)
		}

		r = r+1

	 }
	 return maxLen
}

//[2,5,1,10,10] k=14
//Shrinking taking O(N) time
//Why shrink less than length 3
//when maxlen=3 and sum>k do we need to shrink till <maxLen --> making sure we keep the maxLen as intact. 
//Solution will not working when we have to return subarray, just only len of subarray
//so instead of inner shrinking loop

func LongestSubarraySum_Pattern_Optimal(arr []int, k int)int{
     n:=len(arr)

	 maxLen:=0
    
	 l,r:=0,0
     sum:=0

	 for r<n{
        sum+= arr[r]

		//shrink
		if sum>k{
			sum-=arr[l]
			l++
		}
          
		if sum<=k{
           maxLen = max(maxLen,r-l+1)
		}

		r = r+1

	 }
	 return maxLen
}