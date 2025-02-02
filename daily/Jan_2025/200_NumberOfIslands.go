package daily

// grid := [][]byte{
// 	{'1', '1', '0', '0', '0'},
// 	{'1', '1', '0', '0', '0'},
// 	{'0', '0', '1', '0', '0'},
// 	{'0', '0', '0', '1', '1'},
// }

func NumIslands_bfs(grid [][]byte) int {
	islands := 0
	m, n := len(grid), len(grid[0])

	type pair struct {
		row int
		col int
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				queue := []pair{}
				queue = append(queue, pair{row: i, col: j})

				for len(queue) > 0 {
					currentPair := queue[0]
					queue = queue[1:]

					row, col := currentPair.row, currentPair.col
					grid[row][col] = '0'

					//up
					if row-1 >= 0 && grid[row-1][col] == '1' {
						queue = append(queue, pair{row: row - 1, col: col})
					}

					//down
					if row+1 < m && grid[row+1][col] == '1' {
						queue = append(queue, pair{row: row + 1, col: col})
					}

					//left
					if col-1 >= 0 && grid[row][col-1] == '1' {
						queue = append(queue, pair{row: row, col: col - 1})
					}

					//right
					if col+1 < n && grid[row][col+1] == '1' {
						queue = append(queue, pair{row: row, col: col + 1})
					}
				}

				islands++
			}
		}
	}

	return islands
}

func NumIslands_dfs(grid [][]byte) int {
	islands := 0
	m, n := len(grid), len(grid[0])

	type pair struct {
		row int
		col int
	}

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	visited := make(map[pair]bool)

	var exploreRegion func(int, int)

	exploreRegion = func(i, j int) {
		if i < 0 || j < 0 || i > m-1 || j > n-1 || grid[i][j] == '0' || visited[pair{row: i, col: j}] {
			return
		}

		visited[pair{row: i, col: j}] = true

		for _, d := range directions {
			exploreRegion(i+d[0], j+d[1])
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[pair{row: i, col: j}] {
				exploreRegion(i, j)
				islands++
			}
		}
	}

	return islands
}
