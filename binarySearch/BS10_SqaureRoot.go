package binarySearch

func SqureRoot_Brute(num int) int {
	ans := 1

	for i := 1; i <= num/2; i++ {
		if i*i <= num {
			ans = i
		} else {
			break
		}
	}
	return ans
}

func SqureRoot(num int) int {
	low, high := 1, num

	ans := 1

	for low <= high {
		mid := low + (high-low)/2

		if mid*mid <= num {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return ans
}

// Once low>high --> high will hold the answer
func SqureRoot_Optimal(num int) int {
	low, high := 1, num

	for low <= high {
		mid := low + (high-low)/2

		if mid*mid <= num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return high
}