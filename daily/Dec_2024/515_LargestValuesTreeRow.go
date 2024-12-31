package daily

import "math"

func LargestValues(root *TreeNode) []int {
	ans := []int{}

	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		sz := len(queue)
		max := math.MinInt
		for i := 0; i < sz; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Val > max {
				max = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, max)
	}

	return ans
}
