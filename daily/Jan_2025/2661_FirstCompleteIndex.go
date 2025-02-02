package daily

// mat := [][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}}
// arr := []int{2, 8, 7, 4, 1, 3, 5, 6, 9}

func FirstCompleteIndex(arr []int, mat [][]int) int {
	type pair struct {
		row int
		col int
	}

	n, m := len(mat), len(mat[0])

	matrixMp := make(map[int]pair)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matrixMp[mat[i][j]] = pair{row: i, col: j}
		}
	}

	rowCount := make([]int, n)
	colCount := make([]int, m)

	for i := 0; i < len(arr); i++ {
		pair := matrixMp[arr[i]]

		rowCount[pair.row]++
		colCount[pair.col]++

		if rowCount[pair.row] == m || colCount[pair.col] == n {
			return i
		}

	}

	return -1

}
