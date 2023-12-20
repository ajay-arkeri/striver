package binarySearch

import (
	"math"
)

//TC --> O(max(array) * n)
func KokoBanana_Brute(piles []int, h int) int {
    
	n:= maxInArray(piles)

	for i:=1;i<=n;i++{
		time:= countTime(piles,i)
        // fmt.Println(time)
		if time>h{
			continue
		}

		if time<=h{
			return i
		}

	}

	return -1
}

func countTime(piles []int, n int)int  {
	time:= 0
    
	for i:=0;i<len(piles);i++{
		time+= piles[i]/n

		if piles[i]%n!=0{
			time+=1
		}
	}

	return time
}

func maxInArray(arr []int) int {

	ans := math.MinInt

	for i:=0;i<len(arr);i++{
		if arr[i]>ans{
			ans = arr[i]
		}
	}

	return ans
}


// Since you are having a range 1 to 11 --> can apply binary search

//TC --> O(log2(maxarray) * n)
func KokoBanana(piles []int, h int)int  {
	low,high:= 1,maxInArray(piles)

	for low<=high{
		mid:= low+(high-low)/2

		time:= countTime(piles,mid)

		if time>h{
			low = mid+1
		}else{
			high = mid -1
		}
	}

	return low
}