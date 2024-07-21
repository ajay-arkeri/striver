package binarytrees

func ZigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}

	if root == nil {
		return ans
	}

	flag := true

	queue := []*TreeNode{}

	queue = append(queue, root)

	for len(queue) > 0 {
		n := len(queue)
		index := 0
		temp := make([]int, n)

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			index = i
			if !flag {
				index = n - i - 1
			}

			temp[index] = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, temp)
		flag = !flag
	}
	return ans
}