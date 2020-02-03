package main

import "fmt"

// Time: O(n) (where n is the number of digits, but n is irrelevant since it can't be larger than 10)
// Space: O(n)
// It's probably much simpler to just use digits. I used []byte only because of risk of exceeding int boundaries
func reverse(x int) int {
	if x == 0 {
		return 0
	}

	// Deal with negativeness later
	negative := x < 0
	if negative {
		x = -x
	}

	// Construct string representation of digits
	stringDigits := []byte{}
	for x > 0 {
		digit := x % 10
		stringDigits = append(stringDigits, byte('0'+digit))
		x /= 10
	}

	// Remove leading zeroes
	firstNonZeroIndex := 0
	var digit byte
	for firstNonZeroIndex, digit = range stringDigits {
		if digit != '0' {
			break
		}
	}
	if firstNonZeroIndex == len(stringDigits) {
		return 0
	}
	stringDigitsWithoutLeadingZeroes := stringDigits[firstNonZeroIndex:]

	// Return zero if exceeding limits
	if len(stringDigitsWithoutLeadingZeroes) > len("2147483648") {
		return 0
	}
	if len(stringDigitsWithoutLeadingZeroes) == len("2147483648") &&
		((negative && string(stringDigitsWithoutLeadingZeroes) > "2147483647") ||
			(!negative && string(stringDigitsWithoutLeadingZeroes) > "2147483646")) {
		return 0
	}

	// Add minus if negative and construct the final string
	output := btod(stringDigitsWithoutLeadingZeroes)
	if negative {
		output = -output
	}

	return output
}

func btod(bs []byte) int {
	output := 0
	if len(bs) == 0 {
		return 0
	}
	multi := 1
	for i := len(bs) - 1; i >= 0; i-- {
		output += int(bs[i]-'0') * multi
		multi *= 10
	}
	return output
}

func main() {
	ts := []struct {
		input    int
		expected int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{0, 0},
		{-0, 0},
		{2147483647, 0},
		{-2147483648, 0},
		{7463847412, 0},
		{6463847412, 2147483646},
		{-8463847412, 0},
		{-7463847412, -2147483647},
	}
	for _, tc := range ts {
		actual := reverse(tc.input)
		if tc.expected != actual {
			fmt.Printf("Expected %v but got %v\n", tc.expected, actual)
		}
	}
}
