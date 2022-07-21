package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Time: O(1440 + n) = O(n)
// Space: O(1440) = O(1)
//
// The key insight here is that an array with 1440 entries is constant space, and traversing it once is constant time.
// So one can sort the minutes in linear time, by bucketing them in 1440 buckets (i.e. convert each one to military
// time minutes). Then go once over the minutes and calculate deltas, keeping a minimum. Note that at the end we have
// to check the distance from the latest minute to the first minute as well.
func findMinDifference(timePoints []string) int {
	// No point doing all of this if there are 0 or 1 timepoints!
	if len(timePoints) < 2 {
		return 0
	}

	minExists := make([]bool, 60*24)
	for _, timepoint := range timePoints {
		var (
			parts      = strings.Split(timepoint, ":")
			hours, _   = strconv.Atoi(parts[0])
			minutes, _ = strconv.Atoi(parts[1])
			totalMins  = hours*60 + minutes
		)
		// Optimisation: if we already found this minute before, we can exit early!
		if minExists[totalMins] {
			return 0
		}
		// Bucket sort!
		minExists[totalMins] = true
	}

	var (
		minDistance = math.MaxInt64
		lastMinute  = -1
		firstMinute int
	)
	// This loop is constant time! It looks 1440 time regardless of input size!
	for i, exists := range minExists {
		if !exists {
			continue
		}
		if lastMinute >= 0 {
			minDistance = min(minDistance, i-lastMinute)
		} else {
			firstMinute = i
		}
		lastMinute = i
	}
	// Don't forget to check the last minute against the first minute!
	minDistance = min(minDistance, firstMinute+1440-lastMinute)

	return minDistance
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	ts := []struct {
		input    []string
		expected int
	}{
		{
			input:    []string{"23:59", "00:00"},
			expected: 1,
		},
		{
			input:    []string{"00:00", "23:59", "00:00"},
			expected: 0,
		},
	}
	for _, tc := range ts {
		actual := findMinDifference(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
