package binarytrees

//Maximum number of nodes in a level between any 2 nodes.
//null nodes also counted with width.
type maxWidthPair struct{
	node *TreeNode
	index int
}

//TC - O(N)
//SC ~ O(N)
func WidthOfBinaryTree(root *TreeNode) int {
    ans:=0

    if root==nil{
		return ans
	}

	queue:= []maxWidthPair{}
	queue = append(queue, maxWidthPair{node: root,index: 0})

	for len(queue)>0{
       minIndex:= queue[0].index
       n:=len(queue)

	   left,right:=0,0

	   for i:=0;i<n;i++{
           temp:= queue[0]
		   queue = queue[1:]
           currIndex:= temp.index - minIndex

		   if i==0{
			left = currIndex
		   }

		   if i==n-1{
			right = currIndex
		   }

		   if temp.node.Left!=nil{
			  queue = append(queue, maxWidthPair{node: temp.node.Left,index: 2*currIndex+1})
		   }

		   if temp.node.Right!=nil{
			queue = append(queue, maxWidthPair{node: temp.node.Right,index: 2*currIndex+2})
		   }
	   }

	   ans = max(ans,right-left+1)

	}

	return ans
}