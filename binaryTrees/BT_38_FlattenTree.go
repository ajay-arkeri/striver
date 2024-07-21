package binarytrees

func FlattenTree_Recursion(root *TreeNode) {

	var prev *TreeNode

	var flatten func(*TreeNode)

	flatten = func(node *TreeNode) {

		if node == nil {
			return
		}

		flatten(node.Right)
		flatten(node.Left)

		node.Right = prev
		node.Left = nil
		prev = node

	}

	flatten(root)
}

func FlattenTree_Iterative(root *TreeNode) {

	if root == nil {
		return
	}

	stack := []*TreeNode{}

	stack = append(stack, root)

	for len(stack) > 0 {

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if len(stack) > 0 {
			node.Right = stack[len(stack)-1]
		}

		node.Left = nil

	}
}

//Above 2 solutions took O(N) space
//Using morris traversal to reduce space.

func FlattenTree_Morris(root *TreeNode) {
	curr := root

	for curr != nil {
		if curr.Left != nil {
			prev := curr.Left

			for prev.Right != nil {
				prev = prev.Right
			}

			prev.Right = curr.Right
			curr.Right = curr.Left
			curr.Left = nil
		}

		curr = curr.Right
	}
}