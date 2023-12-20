package binarySearch

//Incomplete --> not from stiver one

func Kth_Largest_SortedArrays_Brute(arr1, arr2 []int, k int) int {

	n1, n2 := len(arr1), len(arr2)
	counter, ans := 0, 0
	i, j := 0, 0

	n := n1 + n2

	for i < n1 && j < n2 {

		if counter == (n - k) {
			break
		}

		if arr1[i] < arr2[j] {
			ans = arr1[i]
			i++
		} else {
			ans = arr2[j]
			j++
		}

		counter++
	}

	if counter != (n - k) {
		if i < n1 {
			ans = arr1[counter-k]
		} else {
			ans = arr2[counter-k]
		}
	}

	return ans

}