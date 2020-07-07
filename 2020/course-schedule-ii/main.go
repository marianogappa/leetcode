package main

import (
	"fmt"
	"reflect"
)

// Time: O(v+e)
// Space: O(v*e) worst case for storing the adjacency list
//
// This is simply topological sort + cycle detection
func findOrder(numCourses int, prerequisites [][]int) []int {
	// In order to run topological sort, constant access to
	// node's edges are necessary.
	adjList := buildAdjacencyList(numCourses, prerequisites)

	// Standard topological search
	result := make([]int, numCourses)
	visited := make([]status, numCourses)
	// Only that the result order is reversed, per exercise
	// requirements.
	n := 0
	for course := range adjList {
		if visited[course] == UNVISITED {
			ok := dfs(course, adjList, visited, result, &n)
			if !ok {
				return []int{}
			}
		}
	}
	return result
}

type status int

const (
	UNVISITED status = iota
	VISITING
	VISITED
)

func dfs(course int, adjList [][]int, visited []status, result []int, n *int) bool {
	if visited[course] == VISITED {
		return true
	}
	// The only difference from standard topological search is
	// that we must abort if a cycle is found.
	if visited[course] == VISITING {
		return false
	}
	visited[course] = VISITING
	for _, prereq := range adjList[course] {
		ok := dfs(prereq, adjList, visited, result, n)
		// Here also aborting if a cycle is found.
		if !ok {
			return false
		}
	}
	visited[course] = VISITED
	result[*n] = course
	(*n)++
	return true
}

func buildAdjacencyList(numCourses int, prerequisites [][]int) [][]int {
	adjList := make([][]int, numCourses)
	for _, prereq := range prerequisites {
		adjList[prereq[0]] = append(adjList[prereq[0]], prereq[1])
	}
	return adjList
}

func main() {
	ts := []struct {
		numCourses    int
		prerequisites [][]int
		expected      []int
	}{
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			expected:      []int{0, 1},
		},
		{
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			expected:      []int{0, 1, 2, 3},
		},
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			expected:      []int{},
		},
	}
	for _, tc := range ts {
		actual := findOrder(tc.numCourses, tc.prerequisites)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.numCourses, tc.prerequisites, tc.expected, actual)
		}
	}
}
