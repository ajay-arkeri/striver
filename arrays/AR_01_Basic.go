package arrays

import (
	"math"
	"sort"
)

//TC - O(NlogN)
func LargestElement_Brute(arr []int) int {
	sort.Ints(arr)
    
	n:=len(arr)

	return arr[n-1]
}

//TC - O(N)
func LargestElement_Optimal(arr []int)int  {
	high:= arr[0]

	for i:=0;i<len(arr);i++{
		if arr[i]>high{
			high = arr[i]
		}
	}

	return high
}

func SecondLargestElementOptimal(arr []int)int  {
	l,sl:=arr[0], math.MinInt


	for i:=0;i<len(arr);i++{
        if arr[i]>l{
			sl = l
			l =arr[i]
		}else if arr[i]>sl{
			sl = arr[i]
		}
	}

	return sl

}

func CheckArrayisSorted(arr []int)bool  {
	n:=len(arr)

	for i:=1;i<n;i++{
       if arr[i]<arr[i-1]{
		return false
	   }
	}

	return true

}

func RemoveDuplicatesInPlace_Brute(arr []int)([]int,int){
	set:=make(map[int]bool)
	unique:= []int{}

    for _,num:= range arr{
		if !set[num]{
			set[num]=true
			unique = append(unique, num)
		}
	}
    
	k:=len(set)
	index:=0

	for _,val:=range unique{
		arr[index]=val
		index++
	}

	return arr,k
}

func RemoveDuplicatesInPlace_Optimal(arr []int)([]int,int){
    n:=len(arr)

	slow:=0
	fast:=1

	for fast<n{
        if arr[fast]!=arr[slow]{
			slow++
			arr[slow] = arr[fast]
		}
		fast++
	}

	return arr,slow+1

}