package binarytrees

type positionNode struct {
	node *TreeNode
	row  int
	col  int
}

func FindBottomLeftValue_Brute(root *TreeNode) int {

	if root == nil {
		return -1
	}
	queue := []positionNode{}
	resultantPair := positionNode{node: root, row: 0, col: 0}
	queue = append(queue, resultantPair)

	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]

		node, row, col := temp.node, temp.row, temp.col

		if resultantPair.row < row {
			resultantPair = temp
		}

		if node.Left != nil {
			queue = append(queue, positionNode{node: node.Left, row: row + 1, col: col - 1})
		}

		if node.Right != nil {
			queue = append(queue, positionNode{node: node.Right, row: row + 1, col: col + 1})
		}
	}

	return resultantPair.node.Val
}

func FindBottomLeftValue(root *TreeNode) int {
	ans := -1

	if root == nil {
		return ans
	}

	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		n := len(queue)

		for i := 0; i < n; i++ {
			temp := queue[i]

			if i == 0 {
				ans = temp.Val
			}

			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}

			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
		}

		queue = queue[n:]
	}

	return ans
}