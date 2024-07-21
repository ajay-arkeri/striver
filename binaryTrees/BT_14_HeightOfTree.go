package binarytrees

//Recursive
func MaxDepth(root *TreeNode) int {
    if root==nil{
        return 0
    }

    lh:= MaxDepth(root.Left)
    rh:=MaxDepth(root.Right)

    return 1+max(lh,rh)
}


//Level Order
func MaxDepth2(root *TreeNode) int {
    ans:= 0

    if root==nil{
        return ans
    }

    queue:= []*TreeNode{}

    queue = append(queue,root)
    
    for len(queue)>0{
        n:=len(queue)

        for i:=0;i<n;i++{
            node:=queue[0]
            queue = queue[1:]

            if node.Left!=nil{
                queue=append(queue,node.Left)
            }

            if node.Right!=nil{
                queue = append(queue, node.Right)
            }
        }

        ans++
    }
    return ans
}
