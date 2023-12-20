package binarySearch

// arr:= []int{1,2,4,6,7,88,99,292}

func BasicBinarySearch_Iterative( arr []int, num int) int {
	low,high:= 0,len(arr)-1
   
	for(low<=high){
		mid:= low + (high-low)/2
       
		if arr[mid]==num{
			return mid
		}

		if arr[mid]<num{
			low = mid+1
		}else{
			high = mid-1
		}
	}
	return -1
}

func BasicBinarySearch_Recursion(arr []int, target ,low, high int )int {
   //base condition
    if low>high{
		return -1
	}

	mid:= low + (high-low)/2

	if arr[mid]==target{
		return mid
	}
    
	if(target>arr[mid]){
		return BasicBinarySearch_Recursion(arr,target,mid+1,high)
	}else{
		return BasicBinarySearch_Recursion(arr,target,low,mid-1)
	}
}