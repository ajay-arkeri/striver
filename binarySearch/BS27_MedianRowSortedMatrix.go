package binarySearch

import (
	"math"
	"sort"
)

//TC --> O(m*n) + (m*n)log(m*n)
func MedianRowSortedMatrix_Brute(mat [][]int) int {
	m, n := len(mat), len(mat[0])

	arr := []int{}

	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			arr = append(arr, mat[i][j])
		}
	}

	sort.Ints(arr)

	return arr[(m*n)/2]
}

// TC --> log2(10^9) * m * log2(n)
func MedianRowSortedMatrix_Optimal(mat [][]int) int {
	m, n := len(mat), len(mat[0])
    req:=(m*n)/2
	low,high:=math.MaxInt,math.MinInt

	for i:=0;i<m;i++{
		if mat[i][0]<low{
			low = mat[i][0]
		}

		if mat[i][n-1]>high{
			high = mat[i][n-1]
		}
	}
    
    
	for low<=high{
		mid := (low+high)/2
       
		count:= elementsLesserThanEqual(mat,mid)
		if count<=req{
			low = mid +1 
		}else{
			high = mid -1 
		}
	    
	}
  
	return low
}


func elementsLesserThanEqual(mat [][]int, target int)int{
	count:=0
	for i:=0;i<len(mat);i++{
		count+= findUpperBound(mat[i],target)
	}

	return count
}

func findUpperBound(arr []int, target int) int {
	low,high:= 0,len(arr)-1
    ans:=len(arr)

	for low<=high{
		mid:=(low+high)/2

		if arr[mid]>target{
		   ans = mid
           high = mid-1
		}else{
           low = mid+1
		}
	}

	return ans
} 