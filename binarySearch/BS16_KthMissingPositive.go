package binarySearch

import "fmt"

//Iterate --> if u encounter a number less than k then k shifts by 1
func FindKthPositive_Brute(arr []int, k int) int {
    for i:=0;i<len(arr);i++{
		if arr[i]<=k{
			k+=1
		}
	}
	return k
}

//Optimal solution
// BS in answers not possible because --> In that u find min, max where some part is possible and some part not possible.
func FindKthPositive(arr []int, k int) int {
    n:= len(arr)

	l,h:= 0,n-1

	for(l<=h){
        mid:= l+(h-l)/2
        
		fmt.Println(arr[mid])
		missing:= arr[mid] - mid + 1
		
		if missing<=k{
			l = mid+1
		}else{
			h = mid-1
		}
	}

	more:=k-(arr[h]- (h + 1))
	fmt.Println(more)
	return arr[h] + more
}