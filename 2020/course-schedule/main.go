package main

import "fmt"

type state int

const (
	unvisited state = iota
	visiting
	visited
)

// Time: O(v + e) where v = vertices and e = edges
// Space: O(v + e)
func canFinish(numCourses int, prerequisites [][]int) bool {
	// Build adjacency list from _prerequisites_.
	adjList := map[int][]int{}
	for _, p := range prerequisites {
		adjList[p[0]] = append(adjList[p[0]], p[1])
	}

	nodeStates := make([]state, numCourses)

	// Check the nodeStates for each course for cycles.
	for i := 0; i < numCourses; i++ {
		if hasCycle(i, adjList, nodeStates) {
			return false
		}
	}
	return true
}

// The key to finding a cycle is: do a dfs traversal of the graph,
// and mark the nodes in the traversal stack as "visiting". Then,
// if you arrive at a "visiting" node while traversing, you have
// a cycle.
func hasCycle(i int, adjList map[int][]int, nodeStates []state) bool {
	switch nodeStates[i] {
	case visited:
		return false
	case visiting:
		return true
	default:
		nodeStates[i] = visiting
		for _, neighbor := range adjList[i] {
			if hasCycle(neighbor, adjList, nodeStates) {
				return true
			}
		}
		nodeStates[i] = visited
		return false
	}
}

func main() {
	ts := []struct {
		numCourses    int
		prerequisites [][]int
		expected      bool
	}{
		{
			numCourses:    2,
			prerequisites: [][]int{{0, 1}},
			expected:      true,
		},
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			expected:      false,
		},
	}
	for _, tc := range ts {
		actual := canFinish(tc.numCourses, tc.prerequisites)
		if tc.expected != actual {
			fmt.Printf("For numCourses: %v, prerequisites: %v expected %v but got %v\n", tc.numCourses, tc.prerequisites, tc.expected, actual)
		}
	}
}
