package arrays

import "math"

//TC--> O(N^2)
func LeadersInArray_Brute(nums []int) []int {
	n := len(nums)
	ans := []int{}
	for i := 0; i < n; i++ {
		leader := true
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] {
				leader = false
				break
			}
		}
		if leader {
			ans = append(ans, nums[i])
		}
	}

	return ans
}

func LeadersInArray_Optimal(nums []int) []int {
    n:=len(nums)
    
	max:=math.MinInt
  
	ans:= []int{}

	for i:=n-1;i>=0;i--{
		if nums[i]>max{
			ans = append(ans, nums[i])
			max = nums[i] 
		}
	}

	return ans
}