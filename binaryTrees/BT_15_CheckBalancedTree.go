package binarytrees

import "math"

//Balanced --> For every node difference between height of left and right subtree <=1
func IsBalanced(root *TreeNode) bool {
     return helper_isBalanced(root)!=-1
}

func helper_isBalanced(root *TreeNode)int{
    if root==nil{
        return 0
    }

    lh:=helper_isBalanced(root.Left)
    if lh==-1{
        return -1
    }

    rh:=helper_isBalanced(root.Right)
    if rh==-1{
        return -1
    }

    if int(math.Abs(float64(lh-rh)))>1{
        return -1
    }

    return 1+max(lh,rh)
}