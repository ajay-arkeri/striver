package arrays

func Rotate_Brute(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	newMatrix := make([][]int, m)

	for i := range newMatrix {
		newMatrix[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			newMatrix[j][n-1-i] = matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		copy(matrix[i], newMatrix[i])
	}
}

func Rotate_Optimal(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	//Transpose
	for i := 0; i < m; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	//Reverse
	for i := 0; i < m; i++ {
		p1, p2 := 0, n-1

		for p1 < p2 {
			matrix[i][p1], matrix[i][p2] = matrix[i][p2], matrix[i][p1]

			p1++
			p2--
		}
	}

}