package sorting

//select the minimum and swap

func SelectionSort(arr []int) []int {

	n:=len(arr)
   
	for i:=0;i<n-1;i++{
        
		min := i 
		for j:=i;j<n;j++{
           if arr[j]<arr[min]{
			min = j
		   }    
		}

		if i!=min{
			arr[i],arr[min] = arr[min],arr[i]
		}
	}

	return arr
}

//TC --> O(N^2)