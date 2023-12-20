package binarySearch

import "fmt"

func PeakElement_Brute(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			if nums[0] > nums[1] {
				return nums[0]
			}
		} else if i == n-1 {
			if nums[n-1] > nums[n-2] {
				return nums[n-1]
			}
		} else {
			if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
				return nums[i]
			}
		}
	}
	return -1
}

func PeakElement_Brute2(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {

		if (i == 0 || nums[i] > nums[i-1]) && (i == n-1 || nums[i] > nums[i+1]) {
			return nums[i]
		}
	}
	return -1
}

//TC --> O(N)
//SC --> O(1)

// Assume array has 1 peak
// Increasing curve --> peak on right
// Decreasing curve --> peak on left
//shrinked to handle overflow cases i==0 and i==n-1
func PeakElement1(nums []int) int {
	n := len(nums)

	if n == 1 {
		return 0
	}

	if nums[0] > nums[1] {
		return 0
	} else if nums[n-1] > nums[n-2] {
		return n - 1
	}

	low, high := 1, n-2

	for low <= high {
		mid := low + (high-low)/2

		fmt.Println(mid)

		//is this peak
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		}

		//increasing and decreasing peak
		if nums[mid] >= nums[mid-1] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

//Assume array has >1 peaks  --> Same above code.

//TC-- O(log2N)