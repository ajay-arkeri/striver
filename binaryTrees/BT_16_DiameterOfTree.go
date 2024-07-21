package binarytrees

func DiameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	helper_diameter(root, &ans)
	return ans
}

func helper_diameter(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}

	lh := helper_diameter(root.Left, ans)
	rh := helper_diameter(root.Right, ans)

	*ans = max(*ans, lh+rh)

	return 1 + max(lh, rh)
}