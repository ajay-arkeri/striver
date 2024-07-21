package binarytrees

//Recursive Solution
func PreorderTraversal_Recursion(root *TreeNode) []int {
	ans := []int{}
	preorderTraversal_Recursive(root, &ans)
	return ans
}

func preorderTraversal_Recursive(node *TreeNode, ans *[]int) {
	if node == nil {
		return
	}

	*ans = append(*ans, node.Val)
	preorderTraversal_Recursive(node.Left, ans)
	preorderTraversal_Recursive(node.Right, ans)
}


//Iterative Solution
func PreorderTraversal_Iterative(node *TreeNode)[]int{
	stack:=[]*TreeNode{} 
	ans:=[]int{}
    
	stack = append(stack, node)

	for len(stack)>0{
       node:= stack[len(stack)-1]
       stack = stack[:len(stack)-1]
   
	   ans = append(ans, node.Val)

	   if node.Right!=nil{
		 stack = append(stack, node.Right)
	   }

	   if node.Left!=nil{
		stack = append(stack, node.Left)
	   }
	}

	return ans
}