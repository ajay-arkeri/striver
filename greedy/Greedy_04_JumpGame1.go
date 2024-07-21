package greedy

//[2,3,1,1,4]
//[3,2,1,0,4]

func CanJump(nums []int) bool {
	n:= len(nums)

	maxJumpIndex:=0 

	for i:=0;i<n;i++{
		
		if i > maxJumpIndex{
			return false
		}

		maxJumpIndex = max(maxJumpIndex,i+nums[i])
	}

	return true
}