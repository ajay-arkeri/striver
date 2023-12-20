package binarySearch

import (
	"math"
)

//Unique elements
//[3,4,5,1,2]
func HowManyTimes_ArrayRotated_Brute(nums []int) int {
	count := 0
	prev := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] < prev {
			return count
		}
		count++
	}

	if count == len(nums) {
		return 0
	}
	return count
}



func  HowManyTimes_ArrayRotated(nums []int)int{
    ans:= math.MaxInt
	index:= 0

	left,right:= 0,len(nums)-1

	for(left<=right){
		mid:= left+(right-left)/2

		if nums[left]<= nums[mid]{
            if nums[left]<ans{
               ans = nums[left]
			   index = left
			}
			left = mid+1
		}else{
			if nums[mid]<ans{
				ans = nums[mid]
				index = mid
			}
			right = mid -1 
		}
	}
   
	return index
}


