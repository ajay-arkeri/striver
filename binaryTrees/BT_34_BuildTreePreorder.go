package binarytrees

//TC --> O(N)
//SC --> O(N) map + O(N) stack space.
func BuildTree_Preorder(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int)
	n := len(inorder)
	for i := 0; i < n; i++ {
		inorderMap[inorder[i]] = i
	}

	return buildTree_Preorder_Helper(preorder, 0, n-1, inorder, 0, n-1, inorderMap)
}

func buildTree_Preorder_Helper(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int, inorderMap map[int]int) *TreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}

	rootVal := preorder[preStart]

	root := &TreeNode{Val: rootVal}
	inRootIndex := inorderMap[rootVal]

	countLeft := inRootIndex - inStart

	root.Left = buildTree_Preorder_Helper(preorder, preStart+1, preStart+countLeft, inorder, inStart, inRootIndex-1, inorderMap)
	root.Right = buildTree_Preorder_Helper(preorder, preStart+countLeft+1, preEnd, inorder, inRootIndex+1, inEnd, inorderMap)
	return root
}