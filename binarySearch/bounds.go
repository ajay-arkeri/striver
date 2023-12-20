package binarySearch

// smallest index such that arr[index]>=target
func LowerBound(arr []int, target int) int {
	ans := len(arr)

	low, high := 0, len(arr)-1

	for low <=high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			ans = mid
			high = mid - 1
		}

	}
	return ans
}

//smallest index such that arr[index]>target
func UpperBound(arr []int, target int)int{
    ans := len(arr)

	low,high:= 0,len(arr)-1

	for(low<=high){
		
		mid:= low+(high-low)/2

		if(arr[mid]>target){
			ans = mid
			high = mid-1
		}else{
			low = mid+1
		}
	}
	return ans
}

//floor --> Largest in array <=x
//ceil ---> Smalled in araray >=x    = lower bound  

func Floor_Ceil(arr []int, target int)(int,int){
    ceil:= LowerBound(arr,target)
	floor:= floor(arr,target)

	return floor,ceil
}

func floor(arr []int, target int)int  {
	ans:=-1

	low,high:= 0,len(arr)-1

	for(low<=high){
       mid:= low+(high-low)/2

	   if arr[mid]==target{
		return mid
	   }else if arr[mid]<target{
		ans = mid
        low=mid+1
	   }else{
		high = mid-1
	   }
	}

	return ans
}

//TC --> O(log2N)