package slidingwindow

//[1,1,1,0,0,0,1,1,1,1,0], k = 2
//TC - O(N^2)
//SC - O(1)
func LongestOnes_Brute(nums []int, k int) int {
    n:=len(nums)
	maxOnes:=0
   
	for i:=0;i<n;i++{
		zeros:=0
		for j:=i;j<n;j++{
           if nums[j]==0{
			zeros++
		   }

		   if zeros>k{
			break	
		   }
		   maxOnes = max(maxOnes,j-i+1)
		}
	}
 

	return maxOnes
}

//TC - O(N + N)
//SC - O(1)
func LongestOnes_Better(nums []int, k int) int {
    n:=len(nums)
	maxOnes:=0
   
	l,r:=0,0
	zeros:=0

	for r<n{
		if nums[r]==0{
			zeros++
		}
          
		//shrink if zeros > k
		for zeros>k{
			if nums[l]==0{
				zeros --
			}
			l++
		}
        maxOnes = max(maxOnes,r-l+1)
		r++
	}

	return maxOnes
}

//TC - O(N)
//SC - O(1)
func LongestOnes_Optimal(nums []int, k int) int {
    n:=len(nums)
	maxOnes:=0
   
	l,r:=0,0
	zeros:=0

	for r<n{
		if nums[r]==0{
			zeros++
		}
          
	    if zeros>k{
			if nums[l]==0{
				zeros--
			}
			l++
		}

		if zeros<=k{
			maxOnes = max(maxOnes,r-l+1)
		}
		r++
	}

	return maxOnes
}