package binarytrees

func RightSideView(root *TreeNode) []int {
	ans := []int{}

	rightViewHelper(root, 0, &ans)

	return ans
}

func rightViewHelper(root *TreeNode, row int, ans *[]int) {
	if root == nil {
		return
	}

	if row == len(*ans) {
		*ans = append(*ans, root.Val)
	}

	rightViewHelper(root.Right, row+1, ans)
	rightViewHelper(root.Left, row+1, ans)
}

func LeftSideView(root *TreeNode) []int {
	ans := []int{}
	leftSideViewHelper(root, 0, &ans)
	return ans
}

func leftSideViewHelper(root *TreeNode, row int, ans *[]int) {
	if root == nil {
		return
	}

	if row == len(*ans) {
		*ans = append(*ans, root.Val)
	}

	leftSideViewHelper(root.Left, row+1, ans)
	leftSideViewHelper(root.Right, row+1, ans)
}