package binarySearch

//Every row in the matrix is sorted

func RowWithMax1s_Brute(mat [][]int) []int {
	
	row:=0
	maxCount:=0

    for i:=0;i<len(mat);i++{
		count:=0
		for j:=0;j<len(mat[i]);j++{
            if mat[i][j]==1{
				count++
			}
		}

		if count>maxCount{
			maxCount = count
			row = i
		}
	}

	return []int{row,maxCount}

}

func RowWithMax1s_Brute_Optimal(mat [][]int) []int {
	
	row:=-1
	maxCount:=0

    for i:=0;i<len(mat);i++{
		count:=0

		low,high:= 0,len(mat[i])-1

		for low<=high{
             mid:=(low+high)/2

			 if mat[i][mid]>0{
				high = mid-1
			 }else{
				low = mid+1
			 }
		}

		count = len(mat[i]) - low


		if count>maxCount{
			maxCount = count
			row = i
		}
	}

	return []int{row,maxCount}

}