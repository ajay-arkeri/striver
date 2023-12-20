package sorting

import "fmt"

func QuickSort(arr []int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr
}

func quickSort(arr []int, low, high int) {
	if low < high {
		partionIndex := partition(arr, low, high)
		fmt.Println(partionIndex)
		quickSort(arr, low, partionIndex-1)
		quickSort(arr, partionIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	i, j := low, high

	pivot := arr[low]

	for i < j {

		for arr[i] <= pivot && i <= j {
			i++
		}

		for arr[j] > pivot && i <= j {
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[j],arr[low] = arr[low],arr[j]

	return j

}