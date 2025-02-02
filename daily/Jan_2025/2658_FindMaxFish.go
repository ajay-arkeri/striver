package daily

// grid := [][]int{{0, 2, 1, 0}, {4, 0, 0, 3}, {1, 0, 0, 4}, {0, 3, 2, 0}}
// grid := [][]int{{0, 5}, {8, 4}}

// 2658
func FindMaxFish_bfs(grid [][]int) int {
	ans := 0
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if grid[i][j] > 0 {
				type pair struct {
					row int
					col int
				}

				queue := []pair{}
				queue = append(queue, pair{row: i, col: j})

				count := 0

				for len(queue) > 0 {
					currentPair := queue[0]
					queue = queue[1:]

					row, col := currentPair.row, currentPair.col

					count += grid[row][col]
					grid[row][col] = 0

					//up
					if row-1 >= 0 && grid[row-1][col] > 0 {
						queue = append(queue, pair{row: row - 1, col: col})
					}
					//down
					if row+1 < m && grid[row+1][col] > 0 {
						queue = append(queue, pair{row: row + 1, col: col})
					}
					//left
					if col-1 >= 0 && grid[row][col-1] > 0 {
						queue = append(queue, pair{row: row, col: col - 1})
					}
					//right
					if col+1 < n && grid[row][col+1] > 0 {
						queue = append(queue, pair{row: row, col: col + 1})
					}
				}

				ans = max(ans, count)
			}

		}
	}

	return ans
}

func FindMaxFish_dfs(grid [][]int) int {
	ans := 0
	m, n := len(grid), len(grid[0])

	var exploreRegion func(int, int) int

	exploreRegion = func(i, j int) int {
		if i < 0 || i > m-1 || j < 0 || j > n-1 || grid[i][j] == 0 {
			return 0
		}

		count := grid[i][j]
		grid[i][j] = 0

		//left
		count += exploreRegion(i, j-1)
		//right
		count += exploreRegion(i, j+1)
		//up
		count += exploreRegion(i-1, j)
		//down
		count += exploreRegion(i+1, j)

		return count
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				ans = max(ans, exploreRegion(i, j))
			}
		}
	}

	return ans
}
