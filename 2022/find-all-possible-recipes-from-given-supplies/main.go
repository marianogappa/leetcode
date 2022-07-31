package main

import (
	"fmt"
	"reflect"
)

// I solved it with DFS + memo intuitively on the first try but it's tricky to answer what the complexity is.
// Also, this exercise is meant to do Kahn's Topological Sort. Which I have to study.
//
// Time:  O(???)
// Space: O(linear to recipes, ingredients & supplies) because the graph will contain recipes & ingredients, plus
//        the isAvailable map.
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	// Available ingredients can be considered "ignorable leaf nodes" in the graph.
	isAvailable := map[string]struct{}{}
	for _, supply := range supplies {
		isAvailable[supply] = struct{}{}
	}

	// Construct a graph from recipes to ingredients (but ingredients can be recipes too)
	graph := map[string]map[string]struct{}{}
	for i, recipe := range recipes {
		// For each recipe, create a node
		graph[recipe] = map[string]struct{}{}
		// For every ingredient of the current recipe...
		for _, ingredient := range ingredients[i] {
			// Ignore ingredients that are available
			if _, ok := isAvailable[ingredient]; ok {
				continue
			}
			// Other ingredients can be (1) recipes, or (2) unknown things that will never be available
			graph[recipe][ingredient] = struct{}{}
		}
	}

	// At this point, we have one node per recipe. Each node can depend on:
	// 1) Nothing                                  -> Doable!
	// 2) At least one ingredient not in the graph -> Not doable!
	// 3) Other recipes                            -> Doable only if recursively resolving ingredients doesn't lead to
	//                                                cycles or ingredients not in the graph.

	// For each recipe, run the check described above
	isRecipeDoable := map[string]bool{}
	for _, recipe := range recipes {
		isRecipeDoable[recipe] = isDoable(recipe, graph, map[string]bool{}, isRecipeDoable)
	}

	// Map -> Slice, because that's what the response type is
	doableRecipes := []string{}
	for recipe, isDoable := range isRecipeDoable {
		if isDoable {
			doableRecipes = append(doableRecipes, recipe)
		}
	}

	return doableRecipes
}

// Recursively resolve recipe dependencies, checking for cycles and with memoisation
func isDoable(recipe string, graph map[string]map[string]struct{}, visited, doableRecipes map[string]bool) bool {
	// No dependencies: Doable!
	if len(graph[recipe]) == 0 {
		return true
	}
	// Found a cycle: Not Doable!
	if visited[recipe] {
		return false
	}
	// Record visited recipes to check for cycles.
	visited[recipe] = true

	// For each requirement:
	for requirement := range graph[recipe] {
		// If the ingredient does not exist: Not Doable!
		if _, ok := graph[requirement]; !ok {
			return false
		}

		// If it exists, it's a recipe, so recursively resolve dependencies, and if not met, Not Doable!
		if !isDoable(requirement, graph, visited, doableRecipes) {
			return false
		}
	}
	// We checked every requirement and they are all ok!
	return true
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
