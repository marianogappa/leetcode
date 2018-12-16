package main

import "fmt"

// Time and space complexity is not useful since there's 32 iterations
func reverseBits(num uint32) uint32 {
	var output uint32
	for i := 31; i >= 0; i-- {
		if isSet(num, uint(i)) {
			output = set(output, uint(31-i))
		}
	}
	return output
}

func isSet(num uint32, i uint) bool {
	return num>>i&1 == 1
}

func set(num uint32, i uint) uint32 {
	return num | (1 << i)
}

func main() {
	ts := []struct {
		input    uint32
		expected uint32
	}{
		{43261596, 964176192},
		{0, 0},
		{1, 2147483648},
	}
	for _, tc := range ts {
		actual := reverseBits(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
