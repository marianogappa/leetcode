package main

import (
	"fmt"
	"strconv"
)

// Time: O(n) one pass over the tokens
// Space: O(n) worst case the whole tokens array will be in the operands stack
//
// This is a very trivial exercise, because the data is in an optimal shape so we don't have to worry about parens or
// anything. All edge cases are taken care of.
//
// If the token is a number, just add it to a stack (here I used an array; a container.List is the right DS though).
// If the token is an operator, pop the last two operands and do the math. Push the result to the stack.
// That's it. When finished, the result is the last pushed value in the stack.
func evalRPN(tokens []string) int {
	operatorFuncs := map[string]func(a, b int) int{
		"+": func(a, b int) int { return a + b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
		"-": func(a, b int) int { return a - b },
	}
	operands := []int{}
	for _, token := range tokens {
		if fn, ok := operatorFuncs[token]; ok {
			operands[len(operands)-2] = fn(operands[len(operands)-2], operands[len(operands)-1])
			operands = operands[:len(operands)-1]
			continue
		}
		n, _ := strconv.Atoi(token)
		operands = append(operands, n)
	}
	return operands[0]
}

func main() {
	ts := []struct {
		input    []string
		expected int
	}{
		{
			input:    []string{"2", "1", "+", "3", "*"},
			expected: 9,
		},
		{
			input:    []string{"4", "13", "5", "/", "+"},
			expected: 6,
		},
		{
			input:    []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			expected: 22,
		},
	}
	for _, tc := range ts {
		actual := evalRPN(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
