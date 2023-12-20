package binarySearch

// TC --> O(m * n )

func FindPeakMatrix_Brute(mat [][]int) []int {
	m, n := len(mat), len(mat[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			b := true

			if i > 0 {
				if mat[i][j] < mat[i-1][j] {
					b = false
				}
			}

			if i < m-1 {
				if mat[i][j] < mat[i+1][j] {
					b = false
				}
			}

			if j > 0 {
				if mat[i][j] < mat[i][j-1] {
					b = false
				}
			}

			if j < n-1 {
				if mat[i][j] < mat[i][j+1] {
					b = false
				}
			}

			if b {
				return []int{i, j}
			}
		}

	}

	return []int{-1, -1}
}


//TC -- O(m * log2n)

func FindPeakMatrix_Optimal(mat [][]int) []int {
    n:=len(mat[0])
    
    low,high:=0,n-1

    for low<=high{
        mid:=(low+high)/2
        row :=  returnMaxElementIndex(mat,mid)

        left:= func()int{
            if mid-1>=0{
                return mat[row][mid-1]
            }else{
                return -1
            }
        }()

        right:= func()int{
            if mid+1<n{
                return mat[row][mid+1]
            }else{
                return -1
            }
        }()
        
        if mat[row][mid]>left && mat[row][mid]>right{
            return []int{row,mid}
        }else if mat[row][mid]<left{
            high = mid -1
        }else{
            low = mid + 1
        }
       

    }

    return []int{-1,-1}
}


func returnMaxElementIndex(mat [][]int,mid int)int{
    
    max:=-1
    index:=-1

    for i:=0;i<len(mat);i++{
        if mat[i][mid]>max{
            max = mat[i][mid]
            index = i
        }
    }

    return index
}