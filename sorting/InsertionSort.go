package sorting

//Takes an element and places it its correct position
//We picked up
// compare and right shift.
func InsertionSort(arr []int)[]int  {
	for i:=0;i<len(arr);i++{
		for j:=i;j>0;j--{
            
			if arr[j]>=arr[j-1]{
                break	
			}else{
                arr[j],arr[j-1] = arr[j-1],arr[j]
			}
		}
	}

	return arr
}