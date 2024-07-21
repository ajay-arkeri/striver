package binarytrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (l *TreeNode) Insert(k int) {
	if l.Val > k {
		//left subtree
		if l.Left == nil {
			l.Left = &TreeNode{Val: k}
		} else {
			l.Left.Insert(k)
		}
	} else {
		//right subtree
		if l.Right == nil {
			l.Right = &TreeNode{Val: k}
		} else {
			l.Right.Insert(k)
		}
	}
}

func GenerateChildrenSumTree() *TreeNode {
	root := &TreeNode{Val: 1}

	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}

	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	return root

}