package binarySearch

//TC --> O(N)
func First_Last_Occurence_Brute(arr []int, target int) (int, int) {
	first, last := -1, -1

	for i := 0; i < len(arr); i++ {
		if arr[i] == target {

			if first == -1 {
				first = i
			}
			last = i
		}
	}

	return first, last
}

//Lower bound = arr[index] >= x
//Upper Bound = arr[index] > x    
//Last occurence = Upper Bound - 1
//First Occurence = Lower Bound 

//TC --> 2 * O(Log2N)
func First_Last_Occurence(arr []int, target int)(int,int){
	//We have check if the number exits and index doesn't point out of the array range --> Edge case
    lb:= LowerBound(arr,target)

	if lb==len(arr) || arr[lb]!=target{
		return -1,-1
	}else{
		return lb,UpperBound(arr,target)-1
	}
}

func First_Last_Occurence_BasicBS(arr []int, target int)(int,int){
    first:= first_occurence(arr,target)

	if first==-1{
		return -1,-1
	}else{
		return first,last_occurence(arr,target)
	}
}

func first_occurence(arr []int, target int)int{
	low,high:= 0,len(arr)-1
    
	first := -1

	for(low<=high){
        mid:= low+(high-low)/2

		if arr[mid]==target{
			first = mid
			high = mid-1
		}else if arr[mid]>target{
            high = mid -1 
		}else{
            low = mid+1
		}
	}

	return first
}

func last_occurence(arr []int, target int)int{
    low,high:= 0,len(arr)-1
    
	last := -1

	for(low<=high){
        mid:= low+(high-low)/2

		if arr[mid]==target{
			last = mid
			low = mid+1
		}else if arr[mid]>target{
            high = mid -1 
		}else{
            low = mid+1
		}
	}

	return last
}

//Total number of occurences question --> last - first + 1