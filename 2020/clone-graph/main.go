package main

type Node struct {
	Val       int
	Neighbors []*Node
}

// Time: O(n)
// Space: O(n)
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// Special case of an empty neighbour list
	rootClonedNode := &Node{Val: node.Val, Neighbors: []*Node{}}
	if len(node.Neighbors) == 0 {
		return rootClonedNode
	}

	// Need to keep a map for constant time access to nodes, in order
	// to link them.
	//
	// Also, a node may be added to the queue many times before it's
	// visited, so we need to flag it as queued. Existence in this
	// map acts as that flag.
	clonedNodes := map[int]*Node{rootClonedNode.Val: rootClonedNode}

	// Finally, a queue to BFS the original nodes. The queue only holds
	// unvisited nodes.
	var queue = []*Node{node}

	// While the queue still holds unvisited nodes...
	for len(queue) > 0 {
		// Pop the queue into a cursor, and consider the node visited.
		curNode := queue[0]
		queue = queue[1:]

		// Clone the node, or get it from the already cloned hashmap.
		clonedNode := clonedNodes[curNode.Val]
		if clonedNode == nil {
			clonedNode = &Node{Val: curNode.Val, Neighbors: []*Node{}}
		}

		// For every neighbour...
		for _, neighbor := range curNode.Neighbors {
			// Clone the neighbour, or get it from the already cloned hashmap.
			if _, ok := clonedNodes[neighbor.Val]; !ok {
				clonedNodes[neighbor.Val] = &Node{Val: neighbor.Val, Neighbors: []*Node{}}
				// Because we're creating this node, it hasn't been queued yet.
				queue = append(queue, neighbor)
			}
			// Add the cloned neighbour to _clonedNode_.
			// Remember that _clonedNode_ is the clone of _curNode_.
			// There won't be double adding, because we only do this for unvisited.
			clonedNode.Neighbors = append(clonedNode.Neighbors, clonedNodes[neighbor.Val])
		}
	}

	return rootClonedNode
}

func main() {
	ns := []*Node{
		nil,
		{1, nil},
		{2, nil},
		{3, nil},
		{4, nil},
	}
	ns[1].Neighbors = []*Node{ns[2], ns[4]}
	ns[2].Neighbors = []*Node{ns[1], ns[3]}
	ns[3].Neighbors = []*Node{ns[2], ns[4]}
	ns[4].Neighbors = []*Node{ns[1], ns[3]}
	cloneGraph(ns[1])
	// ts := []struct {
	// 	input    int
	// 	expected int
	// }{
	// 	{
	// 		input:    1,
	// 		expected: 1,
	// 	},
	// }
	// for _, tc := range ts {
	// 	actual := 1
	// 	if tc.expected != actual {
	// 		fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
	// 	}
	// }
}
