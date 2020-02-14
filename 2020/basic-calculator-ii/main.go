package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Strategy is to split first by +,- and then nested split by *,/.
// Calculate the nested subtotals, and then add or subtract those
// subtotals. All subproblems are linear time, linear space.
//
// Time: O(n)
// Space: O(n)
func calculate(s string) int {
	operands1, operators := splitByRunes(s, []rune{'+', '-'})

	subtotals := []int{}
	for _, operand := range operands1 {
		operandsS, operators := splitByRunes(operand, []rune{'*', '/'})
		operands := intifyOperands(operandsS)
		subtotals = append(subtotals, calculateSubExpressions(operands, operators))
	}

	return calculateSubExpressions(subtotals, operators)
}

// Atois a slice of strings.
// Time: O(n)
// Space: O(n)
func intifyOperands(operandsS []string) []int {
	operands := []int{}
	for _, oS := range operandsS {
		oI, _ := strconv.Atoi(strings.TrimSpace(oS))
		operands = append(operands, oI)
	}
	return operands
}

// Splits a slice of strings by operators, and also gets the operators.
// Time: O(n)
// Space: O(n)
func splitByRunes(s string, rns []rune) ([]string, []byte) {
	operands := strings.FieldsFunc(s, func(r rune) bool {
		for _, rn := range rns {
			if r == rn {
				return true
			}
		}
		return false
	})
	operators := []byte{}
	var l = -1
	for i := 0; i < len(operands)-1; i++ {
		l += len(operands[i]) + 1
		operators = append(operators, s[l])
	}
	return operands, operators
}

// Runs a left-to-right list of calculations (a reduceLeft).
// Time: O(n)
// Space: O(n)
func calculateSubExpressions(operands []int, operators []byte) int {
	if len(operands) == 0 {
		return 0
	}
	var total = operands[0]
	for i := 1; i < len(operands); i++ {
		switch operators[i-1] {
		case '+':
			total += operands[i]
		case '-':
			total -= operands[i]
		case '/':
			total /= operands[i]
		case '*':
			total *= operands[i]
		}
	}
	return total
}

func main() {
	ts := []struct {
		input    string
		expected int
	}{
		{
			input:    "3+2*2",
			expected: 7,
		},
		{
			input:    " 3/2 ",
			expected: 1,
		},
		{
			input:    " 3+5 / 2 ",
			expected: 5,
		},
		{
			input:    "1+1+1",
			expected: 3,
		},
	}
	for _, tc := range ts {
		actual := calculate(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
