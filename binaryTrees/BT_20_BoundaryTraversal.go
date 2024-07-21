package binarytrees

//TC --> O(h) + O(h) + O(N)
//SC --> O(N)
func BoundaryTraversal(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	ans = append(ans, root.Val)
	leftTraversal(root, &ans)
	addLeaves(root, &ans)
	rightTraversal(root, &ans)

	return ans
}

func leftTraversal(node *TreeNode, ans *[]int) {
	cur := node.Left
    
	for cur != nil {
		if !isLeaf(cur) {
			*ans = append(*ans, cur.Val)
		}
		if cur.Left != nil {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
}

func addLeaves(node *TreeNode, ans *[]int) {
	if isLeaf(node) {
		*ans = append(*ans, node.Val)
	}

	if node.Left != nil {
		addLeaves(node.Left, ans)
	}

	if node.Right != nil {
		addLeaves(node.Right, ans)
	}
}

func rightTraversal(node *TreeNode, ans *[]int) {
	cur := node.Right
    temp:=[]int{}

	for cur != nil {
		if !isLeaf(cur) {
			temp = append(temp, cur.Val)
		}

		if cur.Right != nil {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}

	n:=len(temp)

	for i:=n-1;i>=0;i--{
		*ans = append(*ans, temp[i])
	}
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}