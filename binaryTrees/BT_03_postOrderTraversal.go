package binarytrees

//Recursion
func PostorderTraversal_Recursion(root *TreeNode) []int {
	ans := []int{}

	postorderTraversal_recursion(root, &ans)

	return ans
}

func postorderTraversal_recursion(node *TreeNode, ans *[]int) {
	if node == nil {
		return
	}

	postorderTraversal_recursion(node.Left, ans)
	postorderTraversal_recursion(node.Right, ans)
	*ans = append(*ans, node.Val)
}


//Iterative
func PostorderTraversal_Iterative(root *TreeNode) []int {
	 stack1,stack2:=[]*TreeNode{},[]int{}
     if root==nil{
		return stack2
	 }
	 stack1 = append(stack1, root)

	 for len(stack1)>0{
        //pop and put in stack2 
        node:= stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
        stack2 = append(stack2, node.Val)
		//look for left and right, enter into stack1
        
		if node.Left!=nil{
			stack1 = append(stack1, node.Left)
		}

		if node.Right!=nil{
			stack1 = append(stack1, node.Right)
		}
	 }

	 //reverse stack2 and return 

	 for i,j:=0,len(stack2)-1;i<j;i,j = i+1,j-1{
		 stack2[i],stack2[j]=stack2[j],stack2[i]
	 }
     
	 return stack2
	 
}
