package arrays

//TC --> O(2N)
func Sort0s1s2s_Brute(nums []int) []int {
	n := len(nums)

	zeroCount, oneCount, twoCount := 0, 0, 0

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			zeroCount++
		} else if nums[i] == 1 {
			oneCount++
		} else {
			twoCount++
		}
	}

	for i := 0; i < zeroCount; i++ {
		nums[i] = 0
	}

	for i := zeroCount; i < zeroCount+oneCount; i++ {
		nums[i] = 1
	}

	for i := zeroCount + oneCount; i < n; i++ {
		nums[i] = 2
	}

	return nums

}


//Dutch national flag Algo
func Sort0s1s2s_Optimal(nums []int) []int {
	n := len(nums)
    
	low,mid,high:=0,0,n-1


	for mid<=high{
		if nums[mid]==0{
            nums[mid],nums[low] = nums[low],nums[mid]
			low++
			mid++
		}else if nums[mid]==1{
            mid++
		}else{
            nums[mid],nums[high] = nums[high],nums[mid]
			high--	
		}
	}
	

	return nums

}