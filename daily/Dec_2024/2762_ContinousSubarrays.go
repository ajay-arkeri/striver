package daily

import "math"

//arr:=[]int{5,4,2,4}

func ContinuousSubarrays_brute(nums []int) int64 {
	result := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		minimum := math.MaxInt
		maximum := math.MinInt
		for j := i; j < n; j++ {
			minimum = min(nums[j], minimum)
			maximum = max(nums[j], maximum)

			if int(math.Abs(float64(maximum-minimum))) > 2 {
				break
			}
			result++
		}
	}

	return int64(result)
}

func ContinuousSubarrays_better(nums []int) int64 {
	result := 0
	n := len(nums)

	count := make(map[int]int, n)

	l, r := 0, 0

	for r < n {
		count[nums[r]]++

		for l <= r {
			minimum := math.MaxInt
			maximum := math.MinInt
			for val := range count {
				if val < minimum {
					minimum = val
				}

				if val > maximum {
					maximum = val
				}
			}

			if maximum-minimum <= 2 {
				break
			} else {
				count[nums[l]]--
				if count[nums[l]] == 0 {
					delete(count, nums[l])
				}
				l++
			}
		}

		result += r - l + 1
		r++
	}

	return int64(result)
}

func ContinuousSubarrays_optimal(nums []int) int64 {
	result := 0
	n := len(nums)
	minimum := math.MaxInt
	maximum := math.MinInt
	l, r := 0, 0

	for r < n {
		minimum = min(minimum, nums[r])
		maximum = max(maximum, nums[r])

		if maximum-minimum > 2 {
			maximum = nums[r]
			minimum = nums[r]
			l = r

			//move left to find valid window starting point
			for l-1 > 0 && int(math.Abs(float64(nums[l-1]-nums[r]))) <= 2 {
				l--
				maximum = max(maximum, nums[l])
				minimum = min(minimum, nums[l])
			}
		}

		result += r - l + 1
		r++
	}

	return int64(result)
}
