package daily

func FindTargetSumWays(nums []int, target int) int {
	mp := make(map[[2]int]int)
	return calc(nums, target, 0, 0, mp)
}

func calc(nums []int, target int, sum, index int, mp map[[2]int]int) int {
	if index == len(nums) {
		if sum == target {
			return 1
		} else {
			return 0
		}
	}

	key := [2]int{sum, index}
	if val, ok := mp[key]; ok {
		return val
	}

	mp[key] = calc(nums, target, sum+nums[index], index+1, mp) + calc(nums, target, sum-nums[index], index+1, mp)

	return mp[key]
}
