package daily

// edges := [][]int{{0, 2}, {1, 2}, {1, 3}, {2, 4}}
// values := []int{1, 8, 1, 4, 4}
// fmt.Println(maxKDivisibleComponents(5, edges, values, 6))

//Leetcode 2872
func MaxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	adj := map[int][]int{}
	components := 0

	for i := range edges {
		node1, node2 := edges[i][0], edges[i][1]
		adj[node1] = append(adj[node1], node2)
		adj[node2] = append(adj[node2], node1)
	}

	var dfs func(int, int) int

	dfs = func(currentNode, parentNode int) int {
		sum := values[currentNode]

		for _, nextNode := range adj[currentNode] {
			if nextNode != parentNode {
				sum += dfs(nextNode, currentNode)
			}
		}

		if sum%k == 0 {
			components += 1
			return 0
		}

		return sum % k
	}

	dfs(0, -1)

	return components
}
