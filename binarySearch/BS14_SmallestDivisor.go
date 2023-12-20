package binarySearch

import (
	"fmt"
	"math"
)

//find the smallest divisor given a threshold

func SmallestDivisor_Brute(nums []int, threshold int) int {
    max:= findMax(nums)
    
	fmt.Println("max = ",max)
	for i:=1;i<=max;i++{
        if possibleDivisor(nums,i,threshold){
			return i
		}
	}

	return -1
}
//TC --> O(max * n)

func possibleDivisor(nums []int, d int, threshold int)bool  {
	sum:=0

	for i:=0;i<len(nums);i++{
		sum+= int(math.Ceil(float64(nums[i])/float64(d)))
	}

	return sum<= threshold
}

func findMax(arr []int)int{
	max:= math.MinInt

	for i:=0;i<len(arr);i++{
		if arr[i]>max{
			max = arr[i]
		}
	}
	return max
}


func SmallestDivisor(nums []int, threshold int) int {
   min,max:=1,findMax(nums)
  
	ans:=max

	for min<=max{
		 mid:= min+(max-min)/2

        if possibleDivisor(nums,mid,threshold){
					 if mid<ans{
						 ans = mid
					 }
					 max= mid-1
	    	}else{
					min = mid+1
				}
	}

	return ans
}
//TC --> O(log2(max) * n)



// If there is no answer thats within threshold 
// if n>limit {return -1}