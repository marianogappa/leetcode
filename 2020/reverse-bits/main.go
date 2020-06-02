package main

import (
	"fmt"
	"math"
)

// Time: O(1)
// Space: O(1)
func reverseBits(num uint32) uint32 {
	for i := 0; i < 16; i++ {
		l, r := getBit(i, num), getBit(31-i, num)
		if l != r {
			num = setBit(i, num, r)
			num = setBit(31-i, num, l)
		}
	}
	return num
}

func setBit(i int, num uint32, value int) uint32 {
	var mask uint32 = (1 << uint32(i))
	if value == 0 {
		mask = math.MaxUint32 ^ mask
		return num & mask
	}
	return num | mask
}

func getBit(i int, num uint32) int {
	if 1<<uint32(i)&num == 0 {
		return 0
	}
	return 1
}

func main() {
	ts := []struct {
		input    uint32
		expected uint32
	}{
		{
			input:    43261596,
			expected: 964176192,
		},
		{
			input:    4294967293,
			expected: 3221225471,
		},
	}
	for _, tc := range ts {
		actual := reverseBits(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
