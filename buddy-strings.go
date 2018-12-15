package main

import "fmt"

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	diffA, diffB := []byte{}, []byte{}
	var dict [26]int
	hasRepeated := false
	for i := 0; i < len(A); i++ {
		dict[A[i]-'a']++
		if dict[A[i]-'a'] >= 2 {
			hasRepeated = true
		}
		if A[i] != B[i] {
			diffA = append(diffA, A[i])
			diffB = append(diffB, B[i])
		}
		if len(diffA) > 2 {
			return false
		}
	}
	if len(diffA) == 0 {
		return hasRepeated
	}
	if len(diffA) != 2 {
		return false
	}
	if diffA[0] == diffB[1] && diffA[1] == diffB[0] {
		return true
	}
	return false
}

func main() {
	ts := []struct {
		A        string
		B        string
		expected bool
	}{
		{
			"ab",
			"ba",
			true,
		},
		{
			"ab",
			"ab",
			false,
		},
		{
			"aa",
			"aa",
			true,
		},
		{
			"aaaaaaabc",
			"aaaaaaacb",
			true,
		},
		{
			"",
			"aa",
			false,
		},
	}
	for _, tc := range ts {
		actual := buddyStrings(tc.A, tc.B)
		if tc.expected != actual {
			fmt.Printf("For (%v, %v) expected %v but got %v\n", tc.A, tc.B, tc.expected, actual)
		}
	}
}
