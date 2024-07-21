package slidingwindow

//TC - O(N*N)
//SC - O(1)
func NumberOfNiceSubarrays_Brute(nums []int, k int) int {
	n := len(nums)

	count := 0

	for i := 0; i < n; i++ {
		oddCount := 0
		for j := i; j < n; j++ {
			if nums[j]%2 != 0 {
				oddCount++
			}

			if oddCount == k {
				count++
			}
		}
	}

	return count
}


//TC - O(2*2N)
//SC - O(1)
func NumberOfNiceSubarrays_Optimal(nums []int, k int) int {
   x:= countSubarraysWithSumLessThanEqualtok(nums,k)
   y:=countSubarraysWithSumLessThanEqualtok(nums,k-1)

   return x-y
}

func countSubarraysWithSumLessThanEqualtok(nums []int,k int)int{
    if k<0{
        return 0
    }

    n:=len(nums)
    count:=0
    sum:=0

    l:=0
    r:=0

    for r<n{
       sum+=nums[r]%2
       
       for sum>k{
         sum-=nums[l]%2
         l++
       }

       count+=r-l+1
       r++
    }

    return count
}