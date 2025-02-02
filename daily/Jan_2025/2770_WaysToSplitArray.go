package daily

//arr:=[]int{10,4,-8,7}

func WaysToSplitArray(nums []int) int {
	splits := 0

	sum := 0

	for _, val := range nums {
		sum += val
	}

	left := 0

	for i := 0; i < len(nums)-1; i++ {
		left += nums[i]
		sum -= nums[i]

		if left >= sum {
			splits++
		}
	}

	return splits
}
