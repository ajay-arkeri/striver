package arrays

import "fmt"

func SetZeroes_Brute(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				//mark entire row and column as 0

				for k := 0; k < n; k++ {
					if matrix[k][j]!=0{
                       matrix[k][j] = -1
					}
				}

				for k := 0; k < m; k++ {
					if matrix[i][k]!=0{
                       matrix[i][k] = -1
					}
				}
			}
		}
	}

	fmt.Println(matrix)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == -1 {
				matrix[i][j] = 0
			}
		}
	}
}

//O(2*m*n)
//O(m+n)
func SetZeros_Better(matrix [][]int){
	n,m:=len(matrix),len(matrix[0])
    
	row:=make(map[int]bool,n)
    col:= make(map[int]bool,m)
  
	for i:=0;i<n;i++{
		for j:=0;j<m;j++{
			 if matrix[i][j]==0{
				row[i]=true
				col[j] = true
			 }
		}
	}

	for i:=0;i<n;i++{
		for j:=0;j<m;j++{
			if row[i] || col[j]{
				matrix[i][j]=0
			}
		}
	}
}

func SetZeroes_Optimal(matrix [][]int)  {
	n,m:=len(matrix),len(matrix[0])

	//col = matrix[0][...]
	//row = matrix[...][0]
    col0:=1
	
	for i:=0;i<n;i++{
		for j:=0;j<m;j++{
			if matrix[i][j]==0{
                matrix[i][0] = 0
				if j!=0{
                  matrix[0][j] = 0
				}else{
					col0 = 0
				}
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
    
	if matrix[0][0]==0{
		for j:=0;j<m;j++{
			matrix[0][j]=0
		}
	}
	
	if col0==0{
		for i:=0;i<n;i++{
			matrix[i][0]=0
		}
	}
}