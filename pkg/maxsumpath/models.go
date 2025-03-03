package maxsumpath

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// queueItem is a helper struct for BFS, storing the current node and its accumulated path sum.
type queueItem struct {
	node *Node
	sum  int
}
