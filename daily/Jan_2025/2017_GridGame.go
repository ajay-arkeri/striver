package daily

import "math"

// grid := [][]int{{2, 5, 4}, {1, 5, 1}}

func GridGame(grid [][]int) int64 {
	topSum, bottomSum := 0, 0

	for i := 0; i < len(grid[0]); i++ {
		topSum += grid[0][i]
	}

	ans := math.MaxInt64

	for i := 0; i < len(grid[0]); i++ {
		//substract top sum
		topSum -= grid[0][i]
		//find ans

		ans = min(ans, max(topSum, bottomSum))

		//add to bottom sum

		bottomSum += grid[1][i]
	}

	return int64(ans)
}
