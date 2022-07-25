package main

import (
	"fmt"
	"sort"
)

// Time: O(n*k+n*logn) where n is len(indices), and k is length of largest source
// Space: O(n) where n is len(indices), but assuming the output string does not occupy space
//
// This isn't a hard problem to solve at all but one must be careful with off-by-ones.
//
// Initially it looks like a straight up linear time algorithm (considering there are no overlaps),
// as long as one is careful to build a resulting string rather than concatenate linearly every
// time:
//
// Go through the replacements, and either do them or don't do them. For no-go replacements or for
// portions of the string not covered by replacements, copy the original substring.
//
// Note that string equality is linear, and there could be many source[] entries with the same
// index, so the equality check part of the algoritm is linear to len(sources)*len(sources[i]).
//
// The only trick is that we must go through indices in order, and they are not guaranteed to be.
// So we must sort indices, and keep the original order unless we want to do the same to sources
// & targets.
func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	// Sort the indices so we can go through them in order, keeping the original ones
	original_idxs := sortAndReturnIdxs(indices)
	newStr := []byte{}

	stringI := 0
	// For every replacement
	for i, index := range indices {
		// Copy over any substring that is not part of replacements
		if index > stringI {
			toAdd := s[stringI:index]
			newStr = append(newStr, []byte(toAdd)...)
			stringI += len(toAdd)
		}

		// Attempt to make the current replacement onto newStr
		source := sources[original_idxs[i]]
		target := targets[original_idxs[i]]
		// If the source exists in the string...
		if index+len(source) <= len(s) && s[index:index+len(source)] == source {
			// Replace it and advance the string cursor on s
			toAdd := target
			newStr = append(newStr, []byte(toAdd)...)
			stringI += len(source)
		}
	}
	// When all replacements are done, there might still be a suffix left on s
	if stringI < len(s) {
		newStr = append(newStr, []byte(s[stringI:])...)
	}

	return string(newStr)
}

type sortable struct {
	nums, idxs []int
}

func (s sortable) Len() int           { return len(s.nums) }
func (s sortable) Less(i, j int) bool { return s.nums[i] < s.nums[j] }
func (s sortable) Swap(i, j int) {
	s.nums[i], s.nums[j] = s.nums[j], s.nums[i]
	s.idxs[i], s.idxs[j] = s.idxs[j], s.idxs[i]
}

func sortAndReturnIdxs(nums []int) []int {
	idxs := make([]int, len(nums))
	for i := range idxs {
		idxs[i] = i
	}

	sort.Sort(sortable{nums, idxs})

	return idxs
}

func main() {
	ts := []struct {
		s        string
		indices  []int
		sources  []string
		targets  []string
		expected string
	}{
		{
			s:        "abcd",
			indices:  []int{0, 2},
			sources:  []string{"a", "cd"},
			targets:  []string{"eee", "ffff"},
			expected: "eeebffff",
		},
		{
			s:        "abcd",
			indices:  []int{0, 2},
			sources:  []string{"ab", "ec"},
			targets:  []string{"eee", "ffff"},
			expected: "eeecd",
		},
	}
	for _, tc := range ts {
		actual := findReplaceString(tc.s, tc.indices, tc.sources, tc.targets)
		if tc.expected != actual {
			fmt.Printf("For (%v, %v, %v, %v) expected %v but got %v\n", tc.s, tc.indices, tc.sources, tc.targets, tc.expected, actual)
		}
	}
}
