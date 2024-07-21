package binarytrees

import (
	"sort"
)

type nodePostition struct {
	val int
	row int
	col int
}

func VerticalTraversal(root *TreeNode) [][]int {
	ans := [][]int{}

	ds := []nodePostition{}

	preOrderTraversal(root, &ds, 0, 0)

	sort.Slice(ds,func(i, j int) bool {

		if ds[i].col!=ds[j].col{
			return ds[i].col < ds[j].col
		}

		if ds[i].row!=ds[j].row{
			return ds[i].row<ds[j].row
		}
		return ds[i].val<ds[j].val
	})
   
    for i,node:= range ds{
        if i==0 || ds[i-1].col!=ds[i].col{
			ans = append(ans, []int{})
		}

		ans[len(ans)-1] = append(ans[len(ans)-1],node.val)
	}

	return ans
}

func preOrderTraversal(node *TreeNode, ds *[]nodePostition, row, col int) {
	if node == nil {
		return
	}

	*ds = append(*ds, nodePostition{val: node.Val, row: row, col: col})
	preOrderTraversal(node.Left, ds, row+1, col-1)
	preOrderTraversal(node.Right, ds, row+1, col+1)

}