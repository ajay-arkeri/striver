package binarySearch

import (
	"fmt"
	"sort"
)

//Maximum distance between two consequitive stalls is needed to be found.
//arr = []int{0,3,4,7,9,10}  cows  = 4
//ans lies in range of [1,(max-min)+1]  = [1,10]  --> pick each answer and check.  --> at worst case there are 2 cows

//TC --> O(nlogn) + O(max-min) * o(n)
 

func AggressiveCows_brute(arr []int, cows int) int {
    n:= len(arr)
	sort.Ints(arr)
    
	min,max:= arr[0],arr[n-1]
    
	var i int

	for i=1;i<=max-min;i++{
		if canWePlace(arr,cows,i){
			continue
		}else {
			break
		}
	}

	return i-1
}


func canWePlace(arr []int, cows int, d int)bool{
    
	cowsPlaced, lastCowPostition := 1,0

	for i:=1;i<len(arr);i++{
		if arr[i] - arr[lastCowPostition]>=d{
			cowsPlaced++
			lastCowPostition = i
		}
	}

	if cowsPlaced>=cows{
		return true
	}else{
		return false
	}
}

//TC --> O(NlogN) + O(log(max-min)) * log(N)
func AggressiveCows(arr []int,cows int)int{
	n:=len(arr)

	sort.Ints(arr)

    low,high:=1,arr[n-1]


	for low<=high{
		mid:= low+(high-low)/2
        
		fmt.Println(mid)

		if canWePlace(arr,cows,mid){
			low = mid+1
		}else{
			high = mid-1
		}
	}

	return high 
}