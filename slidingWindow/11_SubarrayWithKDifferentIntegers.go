package slidingwindow

//TC - O(N*N)
//SC - O(N)
func SubarraysWithKDistinct_Brute(nums []int, k int) int {
	n := len(nums)

	count := 0

	for i := 0; i < n; i++ {
		mp := make(map[int]int)
		for j := i; j < n; j++ {
			mp[nums[j]]++

			if len(mp) == k {
				count++
			}

			if len(mp) > k {
				break
			}
		}
	}

	return count
}

func SubarraysWithKDistinct_Optimal(nums []int,k int)int{
   x:=findSubarraysWithKDistinct(nums,k)
   y:=findSubarraysWithKDistinct(nums,k-1)

   return x-y
}

func findSubarraysWithKDistinct(nums []int,k int)int{
    n:=len(nums)

	 l,r:=0,0

	 count:=0
     mp:= make(map[int]int)

	 for r<n{
       mp[nums[r]]++

	   //shrink
       for len(mp)>k{
		  mp[nums[l]]--

		  if mp[nums[l]]==0{
			delete(mp,nums[l])
		  }

		  l++
	   }

	   count+= r-l+1

	   r++
	 }

	 return count
}