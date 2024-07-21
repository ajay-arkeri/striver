package arrays

func SpiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])

	top, bottom, left, right := 0, m-1, 0, n-1

	ans := make([]int, 0, m*n)

	for top <= bottom && left <= right {
		//left -> right
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
		top++

		// Top -> bottom

		for i := top; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--

		//Right -> left
		if top <= bottom {
			for i := right; i >= left; i-- {
				ans = append(ans, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			//Bottom -> Top
			for i := bottom; i >= top; i-- {
				ans = append(ans, matrix[i][left])
			}
			left++
		}
	}

	return ans
}