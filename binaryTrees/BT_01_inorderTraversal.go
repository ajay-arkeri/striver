package binarytrees

//Recursive
func InorderTraversal_Recursion(root *TreeNode) []int {
    ans:=[]int{}

    inorderTraversal_recursive(root,&ans)

    return ans
}

func inorderTraversal_recursive(node *TreeNode, ans *[]int){
    if node==nil{
        return 
    }

    inorderTraversal_recursive(node.Left,ans)
    *ans = append(*ans,node.Val)
    inorderTraversal_recursive(node.Right,ans)
}

//Iterative
func InorderTraversal_Iterative(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}

	stack := []*TreeNode{}

	node := root

	for {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			if len(stack) <= 0 {
				break
			}
			node = stack[len(stack)-1]
			ans = append(ans, node.Val)

			stack = stack[:len(stack)-1]
			node = node.Right
		}

	}

	return ans
}