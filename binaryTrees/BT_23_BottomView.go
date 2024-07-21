package binarytrees

import (
	"sort"
)

func BottomView_sort(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}

	mp := make(map[int]int)

	queue := []pair{}
	queue = append(queue, pair{node: root, col: 0})

	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]

		node := temp.node
		dist := temp.col

		mp[dist] = node.Val

		if node.Left != nil {
			queue = append(queue, pair{node: node.Left, col: dist - 1})
		}

		if node.Right != nil {
			queue = append(queue, pair{node: node.Right, col: dist + 1})
		}
	}
   
	keys:= []int{}

    for key:= range mp{
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for _,key:= range keys{
		ans = append(ans, mp[key])
	}

	return ans

}

