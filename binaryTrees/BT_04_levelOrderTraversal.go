package binarytrees

func LevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}

	queue := []*TreeNode{}

	queue = append(queue, root)

	for len(queue) > 0 {
		n := len(queue)
		tempAns := []int{}

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]

			tempAns = append(tempAns, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		ans = append(ans, tempAns)
	}

	return ans
}