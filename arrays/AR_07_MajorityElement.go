package arrays

//Brute --> Take one element at a time scan through entire array and count, if >n/2 then return ans. O(N^2) TC

//Using maps
func MajorityElement_Better(nums []int) int {
    n:=len(nums)

	mp:=make(map[int]int,n)

	for i:=0;i<n;i++{
		mp[nums[i]]++
	}

    for key,val:=range mp{
		if val>n/2{
			return key
		}
	}

	return -1
}

//Moore's Voting algo --> Verification not required if compulsory majority element in array
func MajorityElement_Optimal(nums []int) int {
   n:=len(nums)

   candidate,count:=0,0

   for i:=0;i<n;i++{
	  if count==0{
		candidate = nums[i]
		count =1 
	  }else if nums[i]==candidate{
		count++
	  }else{
		count --
	  }
	}

	//Verification not required if given in question
	count1:=0
	for i:=0;i<n;i++{
		if nums[i]==candidate{
			count1++
		}
	}

	if count1>n/2{
		return candidate
	}

	return -1
}