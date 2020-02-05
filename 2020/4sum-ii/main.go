package main

import "fmt"

// The strategy is to compute hashmaps of the tuples of (A,B) and (C,D),
// both quadratically, and for each tuple on (C,D), see if the negative
// sum exists on the original tuple.
//
// Numbers may be repeated in the slices, so an optimisation is to
// calculate the frequencies to reduce the time complexity.
//
// Note that this solution doesn't leverage the fact that the arrays are
// the same size. There's probably an optimisation behind that.
//
// Time: O(n^2) because calculating tuples of (A,B) and (C,D) is quadratic
// Space: O(n^2) because all tuples of (A,B) are saved
func fourSumCount(A []int, B []int, C []int, D []int) int {
	var (
		numFreqs      = [4]map[int]int{{}, {}, {}, {}} // Unique numbers with frequencies
		pairSumCounts = [2]map[int]int{{}, {}}         // Unique pair sums with freqs
		count         int
	)

	// Calculates the unique numbers (with their freqs) for A, B, C and D
	// Time: O(n)
	// Space: O(n)
	for i, nums := range [][]int{A, B, C, D} {

		for _, num := range nums {
			numFreqs[i][num]++
		}
	}

	// Calculates the unique sums (with their freqs) for (A, B) and (C, D)
	// Time: O(n^2)
	// Space: O(n^2)
	for pairIndex, numsStartIndex := range []int{0, 2} {
		for num1, freq1 := range numFreqs[numsStartIndex] {
			for num2, freq2 := range numFreqs[numsStartIndex+1] {
				pairSumCounts[pairIndex][num1+num2] += freq1 * freq2
			}
		}
	}

	// If the negative of a given pair sum is found in the other map,
	// the frequency multiplication equals the number of ways that
	// numbers can be ordered to produce a zero sum.
	// Time: O(n^2)
	// Space: O(1)
	for sum1, freq1 := range pairSumCounts[0] {
		if freq2, ok := pairSumCounts[1][-sum1]; ok {
			count += freq1 * freq2
		}
	}

	return count
}

func main() {
	ts := []struct {
		a, b, c, d []int
		expected   int
	}{
		{
			a:        []int{1, 2},
			b:        []int{-2, -1},
			c:        []int{-1, 2},
			d:        []int{0, 2},
			expected: 2,
		},
		{
			a:        []int{1, 1},
			b:        []int{1, 1},
			c:        []int{-1, -1},
			d:        []int{-1, -1},
			expected: 16,
		},
	}
	for _, tc := range ts {
		actual := fourSumCount(tc.a, tc.b, tc.c, tc.d)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v,%v,%v) expected %v but got %v\n", tc.a, tc.b, tc.c, tc.d, tc.expected, actual)
		}
	}
}
