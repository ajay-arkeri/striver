package arrays

//Left rotate array by one
func LeftRotatebyOne(arr []int)[]int{
	temp:=arr[0]
	n:=len(arr)

	for i:=1;i<n;i++{
		arr[i-1] = arr[i]
	}

	arr[n-1]=temp

	return arr
}

//TC -> O(d) + O(n-d) + O(d) --> O(n+d)
//SC -> O(d)
func LeftRotateDPlaces_Brute(arr []int,d int)[]int{
	n:=len(arr)
	d = d%n
    temp:= make([]int,d)

	for i:=0;i<d;i++{
		temp[i] = arr[i] 
	}

	for i:=d;i<n;i++{
		arr[i-d] = arr[i]
	}
    
	//j:=0
	for i:=n-d;i<n;i++{
		arr[i] = temp[i - (n-d)]
		//j++
	}

	return arr
}

//TC --> O(d) + O(n-d) + O(n) --> O(2n)
func LeftRotateDPlaces_Optimal(arr []int,d int)[]int{
	reverseArray(arr[:d])
	reverseArray(arr[d:])
	reverseArray(arr)

	return arr
}

func reverseArray(arr []int){
	n:=len(arr)

	for i,j:=0,n-1;i<j;i,j=i+1,j-1{
        arr[i],arr[j] = arr[j],arr[i]
	}
}


func MoveZerosToEnd(arr []int)[]int{
	n:=len(arr)
	j:=0

	for i:=0;i<n;i++{
		if arr[i]!=0{
            arr[i],arr[j] = arr[j],arr[i]
			j++
		}
	}

	return arr
}

func UnionOf2SortedArrays(arr1,arr2 []int)[]int{
	n1,n2:=len(arr1),len(arr2)

	i,j:=0,0

	union:=[]int{}

	for i<n1 && j<n2{
        if arr1[i]<arr2[j]{
           if len(union)==0 || union[len(union)-1]!=arr1[i]{
               union = append(union, arr1[i])
		   }
		   i++
		}else{
           if len(union)==0 || union[len(union)-1]!=arr2[j]{
               union = append(union, arr2[j])
			}
			j++
		}
	}

	for i<n1{
		if len(union)==0 || union[len(union)-1]!=arr1[i]{
               union = append(union, arr1[i])
			   i++
		   }
	}

	for j<n2{
		 if len(union)==0 || union[len(union)-1]!=arr2[j]{
               union = append(union, arr2[j])
			   j++
		   }
	}

	return union
}

func Intersection2SortedArrays(arr1,arr2 []int)[]int  {
    n1,n2:=	len(arr1),len(arr2)
	
	i,j:=0,0

	intersection:=[]int{}

	for i<n1 && j<n2{
       if arr1[i]==arr2[j]{
			intersection = append(intersection, arr1[i])
			i++
			j++
		}else if arr1[i]<arr2[j]{
		i++
	   }else{
         j++
	   }
	}

	return intersection
}