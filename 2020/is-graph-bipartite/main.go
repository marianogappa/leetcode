package main

import (
	"container/list"
	"fmt"
)

// Time: O(v + e) adding all edges to queue (not revisiting); then checking all visited.
// Space: O(v) keeping all nodes in _visited_, plus BFS width in queue.
func isBipartite(graph [][]int) bool {
	if len(graph) <= 1 {
		return true
	}
	visited := make([]int, len(graph))

	for i := range graph {
		// Don't visit nodes twice.
		if visited[i] != 0 {
			continue
		}

		queue := list.New()
		queue.PushFront(i)

		// We "color" the nodes in a BFS traversal by level, 1-2-1-2...
		curGroup := 1

		for queue.Len() > 0 {
			// Clever way of visiting only the next level nodes,
			// even if we add nodes to the queue meanwhile.
			for i := queue.Len(); i > 0; i-- {
				// Pop front node
				headElem := queue.Front()
				queue.Remove(headElem)
				node := headElem.Value.(int)

				// If we already visited this node and it's the wrong "color",
				// then graph is not bipartite.
				if visited[node] != 0 && visited[node] != curGroup {
					return false
				}
				// Don't visit nodes twice.
				if visited[node] != 0 {
					continue
				}

				// "Color" node with current group color.
				visited[node] = curGroup

				// Prepare next BFS traversal level.
				for _, neighbor := range graph[node] {
					queue.PushBack(neighbor)
				}
			}
			// Toggle 1, 2, 1, 2, 1, 2...
			curGroup++
			if curGroup == 3 {
				curGroup = 1
			}
		}
	}

	return true
}

func main() {
	ts := []struct {
		input    [][]int
		expected bool
	}{
		{
			input:    [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}},
			expected: true,
		},
		{
			input:    [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}},
			expected: false,
		},
		{
			input:    [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}, {}},
			expected: true,
		},
		{
			input:    [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}, {0}},
			expected: false,
		},
		{
			input:    [][]int{{}, {}, {}},
			expected: true,
		},
		{
			input:    [][]int{{}},
			expected: true,
		},
		{
			input:    [][]int{{1}, {0}, {4}, {4}, {2, 3}},
			expected: true,
		},
	}
	for _, tc := range ts {
		actual := isBipartite(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
