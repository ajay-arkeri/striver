package binarySearch

// arr := []int{7,8,9,1,2,3,4,5,6}

//All elements are Unique
//Brute --> O(N)

//in normal BS we were used to eliminating one half of array (left/right )
// In this left side not sorted  --> identify the sorted half. at start
//first find the sorted half and then perform check only in the sorted half.

func Rotated_Sorted_Array1(arr []int, target int)int{
    low,high:= 0,len(arr)-1
    
	for(low<=high){
		mid:= low+(high-low)/2
       
        if arr[mid]==target{
			return mid
		}

		//check for sorted side
		
		//left sorted
		if arr[low]<=arr[mid]{
           if target>=arr[low] && target<= arr[mid]{
			  high = mid -1
		   }else{
			  low = mid+1
		   }
		}else {
			if target<=arr[high] && target>= arr[mid]{
				low = mid+1
			}else{
				high = mid-1
			}
		}
	}

	return -1
}




// Array might have duplicates [7,8,1,2,3,3,3,4,5,6] target =3  --> does it exist or not (no returning index)
// problem occurs when arr[low] = arr[mid] = arr[high]  --> need to tackle when this condition occurs [3,1,2,3,3,3,3]
//when the above condition occurs just trim the search space --> low = low +1 and high = high -1
func Rotated_Sorted_Array2(arr []int, target int)int{
  low,high:= 0,len(arr)-1

  for low<=high{
	mid:= low+(high-low)/2
   
	if arr[mid]==target{
		return mid
	}

	if(arr[low]==arr[mid] && arr[high]==arr[mid]){
		low = low+1
		high = high -1
	}else if(arr[low]<=arr[mid]){
		if target>=arr[low] && target<= arr[mid]{
			high = mid -1
		}else{
			low = mid + 1
		}
	}else{
		if target>= arr[mid] && target <= arr[high]{
			low = mid+1
		}else{
			high = mid -1
		}
	}
  }
  return -1
}
//TC --> O(log2N) in most cases
//[3,3,1,3,3,3,3] in worst cases --> O(n/2)