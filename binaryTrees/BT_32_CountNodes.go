package binarytrees

import "math"

//TC --> O(N)
//SC --> stack space O(h) since its a complete binary tree.
func CountNode_Brute(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := CountNode_Brute(root.Left)
	right := CountNode_Brute(root.Right)

	return 1 + left + right
}

//TC --> O(logn)^2 --> at the worst case we traverse log(n) times. we have log(n) for calculating left and right height.
//Sc --> O(log(n))
func CountNodes(root *TreeNode)int{
    if root==nil{
		return 0
	}

	left:=leftHeight(root)
	right:=rightHeight(root)

	if left!=right{
        return 1 + CountNodes(root.Left) + CountNodes(root.Right)
	}else{
		return int(math.Pow(2,float64(left)))-1
	}
}

func leftHeight(root *TreeNode)int{
	height:=0

	for root!=nil{
		root=root.Left
		height++
	}

	return height
}

func rightHeight(root *TreeNode)int{
    height:=0

	for root!=nil{
		root = root.Right
		height++
	}

	return height
}