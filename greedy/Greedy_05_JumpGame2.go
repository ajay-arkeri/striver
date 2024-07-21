package greedy

import (
	"fmt"
	"math"
)

//TC - O(N^N) exponential
func Jump_Recursion(nums []int) int {
	minJumps := math.MaxInt

	recursiveHelper(nums, 0, 0, &minJumps)

	return minJumps
}

func recursiveHelper(nums []int, index int, jumps int, minJumps *int) {
	fmt.Printf("index = %d jumps = %d minJumps = %d \n",index,jumps,*minJumps)
	//base condition
	if index >= len(nums)-1 {
		*minJumps = min(*minJumps, jumps)
		return
	}

	for i := 1 ; i <= nums[index] && index < len(nums); i++ {
		recursiveHelper(nums, index+i, jumps+1, minJumps)
	}
}


//Greedy
//TC - O(N)
//SC - O(1)

func Jump_Greedy(nums []int)int{
    jumps:=0 
    left,right:= 0,0

	for right <len(nums)-1{
       farthest:= 0 

	   for i:=left;i<=right;i++{
		  farthest = max(farthest,nums[i]+i)
	   }

	   left = right+1
	   right = farthest
       jumps++
	}

	return jumps
}