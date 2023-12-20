package binarySearch

import (
	"math"
)

//Unique [4,5,6,7,0,1,2]
//Sorted half may or may not have the answer --> Take the min from sorted half(low/mid) and eliminate the half, we can look in the other half
//If arr[low]<= arr[high] --> then entire space is sorted now so arr[low] is min break the loop --> can avoid BS now
func Minimum_Rotated_Sorted(arr []int)int{
   low,high := 0,len(arr)-1
   
   min:= math.MaxInt
   
   for(low<=high){
	   mid:= low+(high-low)/2
       
	   //already sorted
	   if arr[low]<= arr[high]{
            min = minimum(arr[low],min)
			break
	   }

	   if arr[low]<=arr[mid]{
          min = minimum(arr[low],min)
		  low = mid+1
	   }else{
          min = minimum(arr[mid],min)
          high = mid-1
	   }
   }
   return min
}
//TC --> O(log2N)


//Duplicates
func Minimum_Rotated_Sorted2(arr []int)int{
   low,high := 0,len(arr)-1
   
   min:= math.MaxInt
   
   for(low<=high){
	   mid:= low+(high-low)/2

	   if arr[low]==arr[mid] && arr[mid]==arr[high]{
		  min = minimum(arr[low],min)
		  high--
		  low++
		  continue
	   }
       
	   //already sorted
	   if arr[low]< arr[high]{
            min = minimum(arr[low],min)
			break
	   }

	   if arr[low]<=arr[mid]{
          min = minimum(arr[low],min)
		  low = mid+1
	   }else{
          min = minimum(arr[mid],min)
          high = mid-1
	   }
   }
   return min
}


func minimum(a,b int)int{
	if a<b{
		return a 
	}else {
		return b
	}
}