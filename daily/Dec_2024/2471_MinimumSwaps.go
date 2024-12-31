package daily

func MinimumOperations(root *TreeNode) int {
	queue := []*TreeNode{}
	queue = append(queue, root)
	swaps := 0

	for len(queue) > 0 {
		sz := len(queue)
		arr := []int{}
		for i := 0; i < sz; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			arr = append(arr, node.Val)
		}
		swaps += findSwaps(arr)
	}

	return swaps
}

func findSwaps(arr []int) int {
	l := 0
	swap := 0

	for l < len(arr) {
		currentMinIndex := l
		for i := l + 1; i < len(arr); i++ {
			if arr[i] < arr[currentMinIndex] {
				currentMinIndex = i
			}
		}

		if l != currentMinIndex {
			arr[l], arr[currentMinIndex] = arr[currentMinIndex], arr[l]
			swap++
		}
	}

	l++

	return swap
}
