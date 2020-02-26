package main

import (
	"fmt"
	"reflect"
	"sort"
)

// Union-Find algorithm
type uf struct{ uf []int }

func (u uf) union(a, b int) {
	i, j := u.find(a), u.find(b)
	if i != j {
		u.uf[i] = j
	}
}

func (u uf) find(i int) int {
	if u.uf[i] == 0 {
		return i
	}
	return u.find(u.uf[i])
}

func newUnionFind(size int) uf {
	return uf{make([]int, size+1)}
}

// Time: O(nlogn) in the worst case where all emails belong to the same group
// Space: O(n)
func accountsMerge(accounts [][]string) [][]string {
	// Time: O(n)
	// Space: O(n)
	//
	// The emails map goes from an email address to a pair of ints:
	// [0]: an "AUTO_INCREMENT": 1, 2, 3...
	// [1]: index to one of the accounts that it came from (to grab the name)
	emails := map[string][]int{}
	for i := range accounts {
		for j := 1; j < len(accounts[i]); j++ {
			if _, ok := emails[accounts[i][j]]; !ok {
				emails[accounts[i][j]] = []int{len(emails), i}
			}
		}
	}

	// Time: O(n)
	// Space: O(n)
	//
	// If two emails appear in the same account, we know they belong to the
	// same group so we can `union` them.
	uf := newUnionFind(len(emails))
	for i := range accounts {
		if len(accounts[i]) < 3 {
			continue
		}
		for j := 2; j < len(accounts[i]); j++ {
			uf.union(emails[accounts[i][j-1]][0], emails[accounts[i][j]][0])
		}
	}

	// Time: O(n)
	// Space: O(n)
	//
	// Now we can go through the emails and `find` which group they belong
	// to. We use the `groups` map to go from `group number` to merged's `index`.
	// In here we finally use emails[email][1] to give a name to the merged group.
	groups := map[int]int{}
	merged := [][]string{}
	for email, i := range emails {
		group := uf.find(i[0])
		if _, ok := groups[group]; !ok {
			merged = append(merged, []string{accounts[i[1]][0], email})
			groups[group] = len(merged) - 1
		} else {
			merged[groups[group]] = append(merged[groups[group]], email)
		}
	}

	// Time: O(nlogn)
	// Space: O(1)
	//
	// Emails within each merged group must be sorted.
	for i := range merged {
		sort.Strings(merged[i][1:])
	}

	return merged
}

func main() {
	ts := []struct {
		input    [][]string
		expected [][]string
	}{
		{
			input: [][]string{
				{"John", "johnsmith@mail.com", "john00@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
				{"Mary", "mary@mail.com"},
			},
			expected: [][]string{
				{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"Mary", "mary@mail.com"},
			},
		},
	}
	for _, tc := range ts {
		actual := accountsMerge(tc.input)
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
