package binarytrees

//TC --> O(N) dfs + O(N) BFS worst case
//SC --> O(N)
func DistanceK(root *TreeNode, target *TreeNode, k int) []int {
	ans := []int{}

	if root == nil {
		return ans
	}

	parentMap := make(map[*TreeNode]*TreeNode)

	mapParent(root, &parentMap)

	queue := []*TreeNode{}
	queue = append(queue, target)

	visited := map[*TreeNode]bool{}
	visited[target] = true

	depth := 0

	for len(queue) > 0 {
		n := len(queue)

		if depth == k {
			break
		}

		depth++

		for i := 0; i < n; i++ {
			temp := queue[0]
			queue = queue[1:]

			if temp.Left != nil && !visited[temp.Left] {
				visited[temp.Left] = true
				queue = append(queue, temp.Left)
			}

			if temp.Right != nil && !visited[temp.Right] {
				visited[temp.Right] = true
				queue = append(queue, temp.Right)
			}

			if parent, ok := parentMap[temp]; ok && !visited[parent] {
				visited[parent] = true
				queue = append(queue, parent)
			}
		}

	}

	for len(queue) > 0 {
		ans = append(ans, queue[0].Val)
		queue = queue[1:]
	}

	return ans
}

func mapParent(root *TreeNode, parentMap *map[*TreeNode]*TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		(*parentMap)[root.Left] = root
		mapParent(root.Left, parentMap)
	}

	if root.Right != nil {
		(*parentMap)[root.Right] = root
		mapParent(root.Right, parentMap)
	}
}