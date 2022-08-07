package main

import "fmt"

// Time: O(n)
// Space: O(1)
func canTransform(start string, end string) bool {
	return canTransformRightwards(start, end) || canTransformLeftwards(start, end)
}

func canTransformRightwards(start string, end string) bool {
	strt := []byte(start)
	for i := 0; i < len(strt); i++ {
		if strt[i] == end[i] {
			continue
		}
		if i+1 >= len(strt) {
			return false
		}
		if (string(strt)[i:i+2] == "XL" && end[i] == 'L') || (string(strt)[i:i+2] == "RX" && end[i] == 'X') {
			strt[i], strt[i+1] = strt[i+1], strt[i]
			continue
		}
		return false
	}
	return true
}

func canTransformLeftwards(start string, end string) bool {
	strt := []byte(start)
	for i := len(strt) - 1; i >= 0; i-- {
		if strt[i] == end[i] {
			continue
		}
		if i-1 < 0 {
			return false
		}
		if (string(strt)[i-1:i+1] == "XL" && end[i] == 'X') || (string(strt)[i-1:i+1] == "RX" && end[i] == 'R') {
			strt[i-1], strt[i] = strt[i], strt[i-1]
			continue
		}
		return false
	}
	return true
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
