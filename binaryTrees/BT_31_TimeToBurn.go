package binarytrees

func AmountOfTime(root *TreeNode, start int) int {
	time := 0
	if root == nil {
		return time
	}

	parentMap := make(map[*TreeNode]*TreeNode)

	var startNode *TreeNode

	mapParent_StartNode(root, &parentMap, &startNode, start)

	visited := make(map[*TreeNode]bool)
	visited[startNode] = true

	queue := []*TreeNode{}
	queue = append(queue, startNode)

	for len(queue) > 0 {
		n := len(queue)

		flag := false

		for i := 0; i < n; i++ {
			temp := queue[0]
			queue = queue[1:]

			if temp.Left != nil && !visited[temp.Left] {
				flag = true
				visited[temp.Left] = true
				queue = append(queue, temp.Left)
			}

			if temp.Right != nil && !visited[temp.Right] {
				flag = true
				visited[temp.Right] = true
				queue = append(queue, temp.Right)
			}

			if parent, ok := parentMap[temp]; ok && !visited[parent] {
				flag = true
				visited[parent] = true
				queue = append(queue, parent)
			}
		}
		if flag {
			time++
		}
	}

	return time
}

func mapParent_StartNode(root *TreeNode, parentMap *map[*TreeNode]*TreeNode, startNode **TreeNode, target int) {

	if root.Val == target {
		*startNode = root
	}

	if root.Left != nil {
		(*parentMap)[root.Left] = root
		mapParent_StartNode(root.Left, parentMap, startNode, target)
	}

	if root.Right != nil {
		(*parentMap)[root.Right] = root
		mapParent_StartNode(root.Right, parentMap, startNode, target)
	}
}