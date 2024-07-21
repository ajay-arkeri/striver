package binarytrees

import (
	"fmt"
)

func MaxPathSum(root *TreeNode) int {
	ans := 0
	helper_maxPathSum(root, &ans)
	return ans
}

func helper_maxPathSum(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}

	lSum := helper_maxPathSum(root.Left, ans)
	if lSum < 0 {
		lSum = 0
	}
	rSum := helper_maxPathSum(root.Right, ans)
	if rSum < 0 {
		rSum = 0
	}

	fmt.Println(lSum,rSum)

	*ans = max(*ans, root.Val+root.Val+lSum+rSum)

	fmt.Println(*ans)

	return root.Val + max(lSum, rSum)
}