package main

import "fmt"

// Time: O(logn)
// Space: O(logn)
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	// This is an integer division, which
	// will recurse until n == 0
	temp := myPow(x, n/2)

	// Then we multiply the number by itself
	// logn times (if we were multiplying x
	// we'd multiply n times).
	if n%2 == 0 {
		return temp * temp
	}

	// But if n is odd, we lost 1 in the
	// integer division before. Look:
	//
	// 2/2 = 1
	// 3/2 = 1
	//
	// So remember to add it back (this time
	// we multiply by an extra x; not temp)
	if n > 0 {
		return temp * temp * x
	}

	// And if n is negative, we divide
	// by x instead
	return temp * temp / x
}

func main() {
	ts := []struct {
		x        float64
		n        int
		expected float64
	}{
		{
			x:        2.0,
			n:        10,
			expected: 1024.0,
		},
		{
			x:        2.1,
			n:        3,
			expected: 9.261,
		},
		{
			x:        2.0,
			n:        -2,
			expected: 0.25,
		},
	}
	for _, tc := range ts {
		actual := myPow(tc.x, tc.n)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.x, tc.n, tc.expected, actual)
		}
	}
}
