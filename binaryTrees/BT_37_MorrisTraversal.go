package binarytrees

//threaded binary tree
func MorrisTraversal_Inorder(root *TreeNode)[]int{
    ans:=[]int{}
    
	curr:= root

	for curr!=nil{
        
		if curr.Left==nil{
			ans = append(ans, curr.Val)
			curr = curr.Right
		}else{
            prev:=curr.Left

			for prev.Right!=nil && prev.Right!=curr{
				prev = prev.Right
			}

			if prev.Right==nil{
			   prev.Right = curr
			   curr = curr.Left
			}else{
               prev.Right =nil
			   ans = append(ans, curr.Val)
			   curr = curr.Right
			}

		}


	}


	return ans
}


func MorrisTraversal_PreOrder(root *TreeNode)[]int{
    ans:=[]int{}
    
	curr:= root

	for curr!=nil{
        
		if curr.Left==nil{
			ans = append(ans, curr.Val)
			curr = curr.Right
		}else{
            prev:=curr.Left

			for prev.Right!=nil && prev.Right!=curr{
				prev = prev.Right
			}

			if prev.Right==nil{
			   prev.Right = curr
			   ans = append(ans, curr.Val)
			   curr = curr.Left
			}else{
               prev.Right =nil
			   curr = curr.Right
			}
		}
	}
	return ans
}