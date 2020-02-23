package main

import "fmt"

type state int

const (
	unvisited state = iota
	visiting
	visited
)

// Time: O(n) because the states set will prevent the recursion explosion
// Space: O(n) n states, plus prereqs hash, plus stack up to n
func canFinish(numCourses int, prerequisites [][]int) bool {
	// Prerequisites are not sorted by origin node. If we want to
	// loop through the prereqs of a given node, we have to table
	// scan this array every time. To make this more efficient,
	// construct an array that goes from origin to prereqs.
	var prereqs = make([][]int, numCourses)
	for _, prereq := range prerequisites {
		prereqs[prereq[0]] = append(prereqs[prereq[0]], prereq[1])
	}

	// Nodes can be "unvisited", "visiting" or "visited".
	var states = make([]state, numCourses)

	// Loop through prereqs of each node, and check for cycles.
	for i := 0; i < numCourses; i++ {
		// We only need to check for cycles once, so if while
		// checking for cycles on one node we checked recursively
		// on another, we don't need to do it again.
		if states[i] == unvisited && hasCycle(i, prereqs, states) {
			return false
		}
	}
	return true
}

func hasCycle(i int, prereqs [][]int, states []state) bool {
	// The intermediate "visiting" state allows finding cycles in
	// the recursion. It is legit to arrive to an already visited
	// node: think of a diamond (1 -> 2, 1 -> 3, 2 -> 4, 3 -> 4).
	states[i] = visiting
	for _, prereq := range prereqs[i] {
		switch {
		// The key to detecting the cycle is here. We only recursively
		// visit children, so the only way to have a "visiting" child
		// is with a cycle.
		case states[prereq] == visiting:
			return true
			// A cycle can only happen when we land on a node that has already
			// been visited before. Check for cycles only on those.
		case states[prereq] == visited && hasCycle(prereq, prereqs, states):
			return true
		}
	}
	states[i] = visited
	return false
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
