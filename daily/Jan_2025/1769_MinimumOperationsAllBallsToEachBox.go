package daily

import "math"

func MinOperations_Brute(boxes string) []int {
	n := len(boxes)

	ans := make([]int, n)

	byteArray := []byte(boxes)

	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			if byteArray[j] == '1' {
				sum += int(math.Abs(float64(i - j)))
			}

		}
		ans[i] = sum
	}

	return ans

}

func MinOperations_Optimal(boxes string) []int {
	n := len(boxes)
	result := make([]int, n)

	// Pass from left to right
	moves := 0
	balls := 0
	for i := 0; i < n; i++ {
		result[i] += moves
		if boxes[i] == '1' {
			balls++
		}
		moves += balls
	}

	// Pass from right to left
	moves = 0
	balls = 0
	for i := n - 1; i >= 0; i-- {
		result[i] += moves
		if boxes[i] == '1' {
			balls++
		}
		moves += balls
	}

	return result
}
