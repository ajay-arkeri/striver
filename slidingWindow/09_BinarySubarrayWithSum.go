package slidingwindow

//O(N) time O(N) space solution for count subarrays with sum = k, also works here
//But we can eliminate the extra space for prefix sum since elements are positive only

//TC - O(2 * 2N)
//SC - O(1)
func BinarySubarraysWithSum(nums []int, goal int) int {
     x:= findSubarraysWithGoal(nums,goal)
	 y:=findSubarraysWithGoal(nums,goal-1)
     
	 return x-y
}

func findSubarraysWithGoal(nums []int,goal int)int{
    if goal<0{
		return 0
	}

	n:=len(nums)

	count:=0
	sum:=0 
	l,r:=0,0

	for r<n{
       sum+= nums[r]
        
	   for sum>goal{
		   sum-=nums[l]
		   l++
	   }

       count+= r-l+1

	   r++
	}

	return count
}