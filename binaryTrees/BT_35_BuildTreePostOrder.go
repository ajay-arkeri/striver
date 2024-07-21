package binarytrees

//TC --> O(N) * log(n)
//SC -->O(N)
func BuildTree_Postorder(inorder []int, postorder []int) *TreeNode {
	inorderMap := make(map[int]int)
	n := len(inorder)
	for i := 0; i < n; i++ {
		inorderMap[inorder[i]] = i
	}

	return buildTree_Postorder_Helper(postorder, 0, n-1, inorder, 0, n-1, inorderMap)
}

func buildTree_Postorder_Helper(postorder []int, postStart, postEnd int, inorder []int, inStart, inEnd int, inorderMap map[int]int) *TreeNode {
	if postStart > postEnd || inStart > inEnd {
		return nil
	}

	rootVal := postorder[postEnd]

	root := &TreeNode{Val: rootVal}
	rootIndex := inorderMap[rootVal]
	countLeft := rootIndex - inStart

	root.Left = buildTree_Postorder_Helper(postorder, postStart, postStart+countLeft-1, inorder, inStart, rootIndex-1, inorderMap)
	root.Right = buildTree_Postorder_Helper(postorder, postStart+countLeft, postEnd-1, inorder, rootIndex+1, inEnd, inorderMap)

	return root
}