package binarySearch

import (
	"fmt"
	"math"
)

//Brute --> merging two arrays into a new one
//TC -- O(n1+n2)
//SC -- O(n1+n2)
func MedianSortedArrays_Brute(arr1,arr2 []int)float64{

	n1,n2:=len(arr1),len(arr2)
    i,j:=0,0
	arr3:=[]int{}

	for i<n1 && j<n2{
		if arr1[i]<arr2[j]{
            arr3 = append(arr3, arr1[i])
			i++
		}else{
			arr3 = append(arr3, arr2[j])
			j++
		}
	}

	if i<n1{
		arr3 = append(arr3, arr1[i:]...)
	}
    
	if j<n2{
		arr3 = append(arr3, arr2[j:]...)
	}
    
    n:=len(arr3)

	if n%2==0{
		return float64(arr3[(n/2)-1] + arr3[n/2])/2.0
	}else{
		return float64(arr3[n/2])
	}
}


//Brute better --> no need to store in a 3rd array

func MedianSortedArrays_Better(arr1,arr2 []int)float64{
	n1,n2:=len(arr1),len(arr2)

	n:=n1+n2

	i,j:=0,0
    count:=0

	index1,index2:= (n/2)-1,(n/2)
	index1Val,index2Val:=-1,-1
	fmt.Println("Indxe",index1,index2)

	for i<n1 && j<n2{
        if arr1[i]<arr2[j]{
		   fmt.Println(arr1[i],arr2[j], count)
           if (count==index1) {index1Val = arr1[i]}
		   if (count==index2){index2Val = arr1[i]}
		   count++
		   i++
		}else{
           if (count==index1) {index1Val = arr2[j]}
		   if (count==index2){index2Val = arr2[j]}
		   count++
		   j++
		}

	}
    fmt.Println(count)
    for i<n1{
		if (count==index1) {index1Val = arr1[i]}
		if (count==index2){index2Val = arr1[i]}
		count++
		i++
	}

	for j<n2{
		if (count==index1) {index1Val = arr2[j]}
		if (count==index2){index2Val = arr2[j]}
		count++
		j++
	}
   
	fmt.Println(index1Val,index2Val)

	if n%2==0{
        return float64(index1Val+index2Val)/2.0
	}else{
        return float64(index2Val)
	}
 
}


//Optimal BS approach --> Finding the symmetry
//All the elements on the left should be smaller than all the elements on the right. Cross element check.

func MedianSortedArrays_Optimal(arr1,arr2 []int)float64  {
	
	n1,n2:=len(arr1),len(arr2)
    
	if n1<n2{
		return MedianSortedArrays_Optimal(arr2,arr1)
	}

	n:=n1+n2
	
	low,high:=0,n1

	left:= (n1+n2+1)/2


	for low<=high{
        
		mid1:= (low+high)/2
        mid2:= left-mid1

        l1,l2:= math.MinInt,math.MinInt
        r1,r2:= math.MaxInt,math.MaxInt
        
        if mid1<n1 {r1 = arr1[mid1]}
		if mid2<n2 { r2 = arr2[mid2]}
        if mid1-1>=0 {l1 = arr1[mid1-1]}
		if mid2-1>=0 { l2 = arr2[mid2-1]}

		if l1<=r2 && l2<=r1{
			if n%2==0{
				return float64(max(l1,l2)+min(r1,r2))/2.0
			}

			return float64(max(l1,l2))
		}

        if l1>r2{
			high = mid1-1
		}else{
			low = mid1+1
		}

	}

	return 0.0

}
 