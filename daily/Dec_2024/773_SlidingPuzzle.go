package daily

import (
	"fmt"
	"strings"
)

// board := [][]int{{4, 1, 2}, {5, 0, 3}}

func SlidingPuzzle(board [][]int) int {
	target := "123450"

	directions := [][]int{
		{1, 3},
		{0, 2, 4},
		{1, 5},
		{0, 4},
		{1, 3, 5},
		{2, 4},
	}

	start := boardToString(board)

	visited := map[string]bool{start: true}
	queue := []string{start}
	steps := 0

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			current := queue[0]
			queue = queue[1:]

			if current == target {
				return steps
			}

			zeroIndex := strings.Index(current, "0")

			for _, swapIndex := range directions[zeroIndex] {
				newState := swap(current, zeroIndex, swapIndex)

				if !visited[newState] {
					visited[newState] = true
					queue = append(queue, newState)
				}
			}
		}
		steps++
	}

	return steps
}

func boardToString(board [][]int) string {
	var builder strings.Builder
	for _, row := range board {
		for _, value := range row {
			builder.WriteString(fmt.Sprintf("%d", value))
		}
	}
	return builder.String()
}

func swap(s string, i, j int) string {
	runes := []rune(s)
	runes[i], runes[j] = runes[j], runes[i]
	return string(runes)
}
