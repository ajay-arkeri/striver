package binarySearch

import "math"

func Kth_Element_SortedArrays_Brute(arr1, arr2 []int, k int) int {
	n1, n2 := len(arr1), len(arr2)
	counter, ans := 0, 0
	i, j := 0, 0

	for i < n1 && j < n2 {
		if counter == k {
			break
		}

		if arr1[i] < arr2[j] {
			ans = arr1[i]
			i++
		} else {
			ans = arr2[j]
			j++
		}

		counter++
	}

	if counter != k {
		if i != n1-1 {
			ans = arr1[k-counter]
		} else {
			ans = arr2[k-counter]
		}
	}

	return ans

}

func Kth_Element_SortedArrays_Optimal(arr1, arr2 []int, k int) int {

	n1, n2 := len(arr1), len(arr2)

	low, high := max(0, k-n2), min(n1, k)
	left := k

	for low <= high {
		mid1 := (low + high) / 2
		mid2 := left - mid1

		l1, l2 := math.MinInt,math.MinInt
		r1,r2:= math.MaxInt,math.MaxInt

		if mid1 < n1 {r1 = arr1[mid1]}
		if mid2 < n2 {r2 = arr2[mid2]}
        if mid1-1>=0 {l1 = arr1[mid1-1]}
		if mid2-1>=0 {l2 = arr2[mid2-1]}


		if (l1<=r2 && l2<=r1){
			return max(l1,l2)
		}

		if l1>r2{
			high = mid1-1
		}else{
			low = mid1+1
		}

	}
     
	return 0
}