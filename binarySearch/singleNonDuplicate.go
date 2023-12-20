package binarySearch

//Guranteed single
//1,1,2,3,3,4,4,8,8
func SingleNonDuplicate_Brute(nums []int) int {
   
	n:= len(nums)
    
	if n==1{
		return nums[0]
	}

   for i:=0;i<n;i++{
	  if i==0{
		if nums[0]!=nums[1]{
			return nums[0]
		}
	  }else if i==n-1{
        if nums[n-1]!=nums[n-2]{
			return nums[n-1]
		} 
	  }else{  
		if nums[i]!=nums[i-1] && nums[i]!=nums[i+1]{
			return nums[i]
		}
	  }
   }
   return -1
}


//(even,odd) --> element is in the right half
//(odd,even) ---> element is in the left half
func SingleNonDuplicate(nums []int)int{
	n:=len(nums)
   
	if n==1 {
		return nums[0]
	}

	if nums[0]!=nums[1]{
		return nums[0]
	}
    
	if nums[n-1]!=nums[n-2]{
		return nums[n]
	}
  
    low,high:=1,n-2

	for(low<=high){
		mid:= low+(high-low)/2

		if nums[mid]!=nums[mid-1] && nums[mid]!=nums[mid+1]{
			return nums[mid]
		}

		if( mid%2==1 && nums[mid-1]==nums[mid]) || (mid%2==0 && nums[mid]==nums[mid+1]){
            low = mid+1
		}else{
			high = mid-1
		}
	}

	return -1
}