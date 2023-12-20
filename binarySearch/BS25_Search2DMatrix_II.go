package binarySearch

//Brute --> O(m*n)

//Better - O(m * logm)
func Search2DMatrix_II_Better(matrix [][]int, target int) bool {
       m:=len(matrix)

    for i:=0;i<m;i++{
        low,high:=0,len(matrix[i])-1

        for low<=high{
            mid:=(low+high)/2

            if matrix[i][mid]==target{
                return true
            }else if matrix[i][mid]>target{
                high = mid - 1
            }else{
                low = mid + 1 
            }
        }
    }

    return false
}


// Start from right top/ left bottom corner --> eliminate a row or column based on condition while searching the target.

func Search2DMatrix_II_Optimal(matrix [][]int, target int) bool {
    m,n:=len(matrix),len(matrix[0])

	row,col:= 0,n-1

	for row<m && col>=0{
		if matrix[row][col]==target{
			return true
		}else if matrix[row][col]>target{
			col--
		}else{
			row++
		}
	}
	return false
}