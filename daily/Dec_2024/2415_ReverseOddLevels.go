package daily

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ReverseOddLevels_BFS(root *TreeNode) *TreeNode {
	queue := []TreeNode{}

	level := 0

	for len(queue) > 0 {
		size := len(queue)
		t := []TreeNode{}

		oddLevel := level%2 == 0
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if oddLevel {
				t = append(t, node)
			}

			queue = append(queue, *node.Left)
			queue = append(queue, *node.Right)
		}

		if oddLevel {
			reverseNodeValues(t)
		}
	}

	return root
}

func reverseNodeValues(arr []TreeNode) {
	i, j := 0, len(arr)-1

	for i < j {
		arr[i].Val, arr[j].Val = arr[j].Val, arr[i].Val
	}
}

func ReverseOddLevels_DFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// Start DFS with root's children at level 1
	dfs(root.Left, root.Right, 1)
	return root
}

func dfs(left, right *TreeNode, level int) {
	if left == nil || right == nil {
		return
	}

	// If the level is odd, swap the values
	if level%2 != 0 {
		left.Val, right.Val = right.Val, left.Val
	}

	// Recur for the next level
	dfs(left.Left, right.Right, level+1)
	dfs(left.Right, right.Left, level+1)
}
