package main

import (
	"fmt"
	"reflect"
	"sort"
)

// Time: O(n*n!) because constructing a permutation takes n iterations
// Space: O(n*n!) count of permutations times the size of each permutation
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	// Sorting is necessary to ignore duplicates.
	sort.Ints(nums)

	// Used is used with backtracking to not reuse
	// indices (NOT values!) whilst constructing
	// permutations.
	used := make([]bool, len(nums))

	results := [][]int{}
	partial := make([]int, 0, len(nums))
	backtrack(partial, used, nums, &results)

	return results
}

func backtrack(partial []int, used []bool, nums []int, results *[][]int) {
	// Only push results when all numbers are used in a permutation.
	if len(partial) == len(nums) {
		*results = append(*results, clone(partial))
		return
	}
	// In Subsets, a recursive call starts from "the next index of nums",
	// but in Permutations we go through the array and ignore the used
	// indices.
	for i := 0; i < len(nums); i++ {
		// Ignore already used indices in this permutation.
		if used[i] {
			continue
		}

		// Mark index used, and append it to the partial permutation.
		used[i] = true
		partial = append(partial, nums[i])
		backtrack(partial, used, nums, results)

		// Backtrack.
		partial = partial[:len(partial)-1]
		used[i] = false

		// Ignore all subsequent numbers that are equal to the current.
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
	}
}

func clone(ns []int) []int {
	cloned := make([]int, len(ns))
	copy(cloned, ns)
	return cloned
}

func main() {
	ts := []struct {
		input    []int
		expected [][]int
	}{
		{
			input:    []int{1, 1, 2},
			expected: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
	}
	for _, tc := range ts {
		actual := permuteUnique(tc.input)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
