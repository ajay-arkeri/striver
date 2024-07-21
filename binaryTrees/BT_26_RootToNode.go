package binarytrees

import "strconv"

func RootToNode(root *TreeNode, target int) []int {
	ans := []int{}

	if root == nil {
		return ans
	}

	rootToNodeHelper(root, &ans, target)

	return ans
}

func rootToNodeHelper(node *TreeNode, ans *[]int, target int) bool {
	if node == nil {
		return false
	}

	*ans = append(*ans, node.Val)

	if node.Val == target {
		return true
	}

	if rootToNodeHelper(node.Left, ans, target) || rootToNodeHelper(node.Right, ans, target) {
		return true
	}

	*ans = (*ans)[:len(*ans)-1]

	return false

}

// Leetcode 257
// ["1->2->5","1->3"]
func BinaryTreePaths_RootToLeaf(root *TreeNode) []string {
	ans := []string{}

	if root == nil {
		return ans
	}

	rootToLeafHelper(root, &ans, "")

	return ans
}

func rootToLeafHelper(node *TreeNode, ans *[]string, temp string) {
	if node.Left==nil && node.Right==nil {
		temp+=strconv.Itoa(node.Val)
		*ans = append(*ans, temp)
		return
	}
	
	temp += strconv.Itoa(node.Val)
	temp+="->"

	if node.Left!=nil{
		rootToLeafHelper(node.Left, ans,temp)
	}
	
	if node.Right!=nil{
		rootToLeafHelper(node.Right,ans,temp)
	}
}