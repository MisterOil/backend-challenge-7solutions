package maxsumpath

func BuildTree(data [][]int) *Node {
	if len(data) == 0 {
		return nil
	}

	// Create a parallel 2D slice of Node pointers.
	nodes := make([][]*Node, len(data))
	for i := range data {
		nodes[i] = make([]*Node, len(data[i]))
		for j := range data[i] {
			nodes[i][j] = &Node{Val: data[i][j]}
		}
	}

	// Link each node to its left and right children.
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data[i]); j++ {
			nodes[i][j].Left = nodes[i+1][j]
			nodes[i][j].Right = nodes[i+1][j+1]
		}
	}

	return nodes[0][0]
}

func MaxSumPathBFS(root *Node) int {
	if root == nil {
		return 0
	}

	maxSum := 0
	queue := []queueItem{{node: root, sum: root.Val}}

	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]

		node := current.node
		sumSoFar := current.sum

		if node.Left == nil && node.Right == nil {
			if sumSoFar > maxSum {
				maxSum = sumSoFar
			}
		} else {
			if node.Left != nil {
				queue = append(queue, queueItem{
					node: node.Left,
					sum:  sumSoFar + node.Left.Val,
				})
			}
			if node.Right != nil {
				queue = append(queue, queueItem{
					node: node.Right,
					sum:  sumSoFar + node.Right.Val,
				})
			}
		}
	}
	return maxSum
}
