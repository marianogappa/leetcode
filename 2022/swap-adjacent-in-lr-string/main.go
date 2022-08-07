package main

import (
	"fmt"
	"strings"
)

// Time: O(s + e)
// Space: O(s + e) in the worst case all letters are Ls and Rs (we store all idxs)
//
// Since XL -> LX and RX -> XR, one way to see this is that, in a current of Xs,
// Ls can flow to the left as much as they want, and Rs can flow to the right,
// but they cannot jump on each other.
//
// So the first linear check one can make is that, when stripping out the Xs,
// the strings must be the same.
//
// Guaranteed to be the same, there are an equal number of Ls and Rs, but they
// don't have to be in the same indices.

// The rule should be:
// - For each "L" in "start", it has to be in a <= index on "end".
// - For each "R" in "start", it has to be in a >= index on "end".
//
// That's it. If that's true, you can transform them.
func canTransform(start string, end string) bool {
	if strings.ReplaceAll(start, "X", "") != strings.ReplaceAll(end, "X", "") {
		return false
	}
	startL := idxsOfByte(start, 'L')
	startR := idxsOfByte(start, 'R')
	endL := idxsOfByte(end, 'L')
	endR := idxsOfByte(end, 'R')

	for i := 0; i < len(startL); i++ {
		if startL[i] < endL[i] {
			return false
		}
	}

	for i := 0; i < len(startR); i++ {
		if startR[i] > endR[i] {
			return false
		}
	}

	return true
}

func idxsOfByte(str string, byt byte) []int {
	idxs := []int{}
	for i := 0; i < len(str); i++ {
		if str[i] == byt {
			idxs = append(idxs, i)
		}
	}
	return idxs
}

func main() {
	ts := []struct {
		start, end string
		expected   bool
	}{
		{
			start:    "RXXLRXRXL",
			end:      "XRLXXRRLX",
			expected: true,
		},
		{
			start:    "X",
			end:      "L",
			expected: false,
		},
		{
			start:    "XXXXXLXXXX",
			end:      "LXXXXXXXXX",
			expected: true,
		},
		{
			start:    "XLXRRXXRXX",
			end:      "LXXXXXXRRR",
			expected: true,
		},
	}
	for _, tc := range ts {
		actual := canTransform(tc.start, tc.end)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.start, tc.end, tc.expected, actual)
		}
	}
}
