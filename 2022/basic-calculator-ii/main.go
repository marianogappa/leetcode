package main

import (
	"fmt"
	"strconv"
)

// Time: O(n)
// Space: O(n)
//
// The parsing step seems pretty trivial, since you will always start with a number, then an
// operator, then a number and so on. There will be n operators and n+1 operands.
//
// Only trick is the '-' character, which can be the first character in an operand, or an
// operator in itself.
//
// Also, remember to trim spaces before popping each token.
//
// In terms of the calculation, there's only two levels of precedence, so * & / can be done first,
// and + & - later.
//
// The +/- part is trivial, but for the *// part the trick is to flatten the operands by running
// the operator on them.
func calculate(s string) int {
	operands, operators := parse(s)
	if len(operands) == 1 {
		return operands[0]
	}
	if len(operators) == 0 {
		return 0
	}
	operands, operators = solveMulDiv(operands, operators)
	return solvePlusMinus(operands, operators)
}

func solveMulDiv(operands []int, operators []byte) ([]int, []byte) {
	var (
		newOperands  = []int{operands[0]}
		newOperators []byte
	)
	for i, operator := range operators {
		switch operator {
		case '*':
			newOperands[len(newOperands)-1] *= operands[i+1]
		case '/':
			newOperands[len(newOperands)-1] /= operands[i+1]
		case '+', '-':
			newOperators = append(newOperators, operator)
			newOperands = append(newOperands, operands[i+1])
		}
	}
	return newOperands, newOperators
}

func solvePlusMinus(operands []int, operators []byte) int {
	var runningSum = operands[0]
	for i, operator := range operators {
		if operator == '+' {
			runningSum += operands[i+1]
		} else {
			runningSum -= operands[i+1]
		}
	}
	return runningSum
}

func parse(s string) ([]int, []byte) {
	var (
		isOperator bool
		i, ln      int
		operator   byte
		num        int
		nums       []int
		operators  []byte
	)
	for i < len(s) {
		i += countLeadingSpaces(s[i:])
		if i >= len(s) {
			break
		}
		if isOperator {
			operator, ln = parseOperator(s[i:])
			operators = append(operators, operator)
		} else {
			num, ln = parseNum(s[i:])
			nums = append(nums, num)
		}
		i += ln
		isOperator = !isOperator
	}
	return nums, operators
}

func countLeadingSpaces(s string) (i int) {
	for {
		if i == len(s) || s[i] != ' ' {
			return i
		}
		i++
	}
}

func parseNum(s string) (int, int) {
	var i int
	for i = 0; i < len(s); i++ {
		if !((s[i] >= '0' && s[i] <= '9') || (s[i] == '-' && i == 0)) {
			break
		}
	}
	num, _ := strconv.Atoi(s[:i])
	return num, len(s[:i])
}

func parseOperator(s string) (byte, int) {
	return s[0], 1
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
			input:    " -1 ",
			expected: -1,
		},
		{
			input:    " -1+1 ",
			expected: 0,
		},
		{
			input:    " 0000 ",
			expected: 0,
		},
		{
			input:    "0-2147483647",
			expected: -2147483647,
		},
	}
	for _, tc := range ts {
		actual := calculate(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
