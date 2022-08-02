package main

import (
	"container/list"
	"fmt"
	"reflect"
)

// There are two keys to how to solve this:
//
// 1) Elements form a graph structure, because recipes require ingredients, but there are multiple "root nodes" and
//    potentially nodes can have edges back to the roots.
// 2) The problem is asking for which recipes there are no cycles in the graph, which Topological Sort via Kahn's
//    algorithm will answer.
//
// Time: O(r + s + i) essentially V + E for Kahn's, and a few passes over recipes + supplies to construct structures.
// Space: O(r + s + i) for constructing the graph, the queue & the set
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	// Make a recipe set, so that later on recipes can be efficiently detected, to be added to response.
	recipeSet := map[string]struct{}{}
	for _, recipe := range recipes {
		recipeSet[recipe] = struct{}{}
	}

	// Supplies have no dependencies. They have an "in-degree zero" according to Kahn's algorithm.
	inDegree := map[string]int{}
	queue := list.New()
	for _, supply := range supplies {
		inDegree[supply] = 0
		queue.PushBack(supply)
	}

	// Remember that the graph structure goes from dependency to product, as in: "A -> B", "A is required to get to B".
	// In this case, ingredient leads to recipe. And since the arrow goes "into" recipe, the in-degree of recipe
	// is incremented.
	graph := map[string][]string{}
	for i, recipe := range recipes {
		for _, ingredient := range ingredients[i] {
			graph[ingredient] = append(graph[ingredient], recipe)
			inDegree[recipe]++
		}
	}

	// Run Kahn's algorithm on graph.
	//
	// Note that cycles (e.g. recipe A depends on recipe B, and vice-versa) don't affect this algorithm, because those
	// recipes will never have inDegree == 0, never get pushed into the queue, and thus the algorithm terminates
	// without adding them
	doableRecipes := []string{}
	for queue.Len() > 0 {
		front := queue.Front()
		elem := front.Value.(string)
		queue.Remove(front)

		if _, ok := recipeSet[elem]; ok {
			doableRecipes = append(doableRecipes, elem)
		}

		for _, prod := range graph[elem] {
			inDegree[prod]--
			if inDegree[prod] == 0 {
				queue.PushBack(prod)
			}
		}
	}

	return doableRecipes
}

func main() {
	ts := []struct {
		recipes     []string
		ingredients [][]string
		supplies    []string
		expected    []string
	}{
		{
			recipes:     []string{"bread"},
			ingredients: [][]string{{"yeast", "flour"}},
			supplies:    []string{"yeast", "flour", "corn"},
			expected:    []string{"bread"},
		},
		{
			recipes:     []string{"bread", "sandwich"},
			ingredients: [][]string{{"yeast", "flour"}, {"bread", "meat"}},
			supplies:    []string{"yeast", "flour", "meat"},
			expected:    []string{"bread", "sandwich"},
		},
		{
			recipes:     []string{"bread", "sandwich", "burger"},
			ingredients: [][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}},
			supplies:    []string{"yeast", "flour", "meat"},
			expected:    []string{"bread", "sandwich", "burger"},
		},
	}
	for _, tc := range ts {
		actual := findAllRecipes(tc.recipes, tc.ingredients, tc.supplies)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For (%v,%v,%v) expected %v but got %v\n", tc.recipes, tc.ingredients, tc.supplies, tc.expected, actual)
		}
	}
}
