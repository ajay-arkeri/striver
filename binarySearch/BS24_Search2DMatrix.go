package binarySearch

//TC --> O(N*M)
func Search2DMatrix_Brute(matrix [][]int, target int) bool {

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

//TC - O(N) + O(log2(M))
func Search2DMatrix_Better(matrix [][]int, target int) bool {

	for i := 0; i < len(matrix); i++ {
		end:=len(matrix[i])-1
		if(target>=matrix[i][0] && target<= matrix[i][end]){
            low,high:= 0,end

			for low<=high{
				mid:= (low+high)/2

				if matrix[i][mid]==target{
					return true
				}else if matrix[i][mid]>target{
					high = mid - 1
				}else{
					low = mid + 1
				}
			}

			if matrix[i][low]==target{
				return true
			}
		}
	}
	return false
}

//TC - O(log2(m*n))
func Search2DMatrix_Optimal(matrix [][]int, target int) bool {
    m,n:= len(matrix),len(matrix[0])
	low,high:=0,(m*n)-1


	for low<=high{
        mid:= (low+high)/2

		row,col:= mid/n, mid%n

		if matrix[row][col]==target{
			return true
		}else if matrix[row][col]>target{
            high  = mid -1
		}else{
			low = mid + 1
		}
	}

	return false
}
