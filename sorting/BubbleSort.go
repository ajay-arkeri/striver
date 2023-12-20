package sorting

//Pushes the maximum element to the last
//Worst and average case --> O(N^2)
//Best case --> O(N) --> need some check to see when no swap is done.
func BubbleSort(arr []int)[]int{
    n:=len(arr)

	for i:=0;i<n;i++{
        swap:=false
		for j:=0;j<n-i-1;j++{
            if arr[j]>arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
				swap = true
			}
		}

		if !swap{
			return arr
		}

	}

	return arr
}