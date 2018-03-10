package main

import "fmt"

// Implements this Regex
// ' *[+-]?[0-9]*(\.[0-9]+)?([eE][+-]?[0-9]+)? *'
// Plus ensuring there's at least one number before or after the decimal point
func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	var i = 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i >= len(s) {
		return false
	}
	if s[i] == '+' || s[i] == '-' {
		i++
	}
	if i >= len(s) {
		return false
	}
	var atLeastOne = false
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		atLeastOne = true
		i++
	}
	if i >= len(s) {
		return atLeastOne
	}
	var atLeastOneAfterDecimal = false
	if s[i] == '.' {
		i++
		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
			atLeastOneAfterDecimal = true
			i++
		}
	}
	if i >= len(s) {
		return atLeastOne || atLeastOneAfterDecimal
	}
	if s[i] == 'e' || s[i] == 'E' {
		i++
		if i >= len(s) {
			return false
		}
		if s[i] == '+' || s[i] == '-' {
			i++
		}
		if i >= len(s) {
			return false
		}
		var atLeastOneAfterE = false
		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
			atLeastOneAfterE = true
			i++
		}
		if !atLeastOneAfterE {
			return false
		}
	}
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i != len(s) {
		return false
	}
	return atLeastOne || atLeastOneAfterDecimal
}

func main() {
	var ts = []struct {
		i string
		e bool
	}{
		{i: "0", e: true},
		{i: " 0.1 ", e: true},
		{i: "abc", e: false},
		{i: "1 a", e: false},
		{i: "2e10", e: true},
		{i: "-1", e: true},
		{i: "-+1", e: false},
		{i: "6.022E23", e: true},
		{i: "   6.022E23  ", e: true},
		{i: "   +6.022E23  ", e: true},
		{i: "   -6.022E23  ", e: true},
		{i: "   - 6.022E23  ", e: false},
		{i: "   - 6.022e23  ", e: false},
		{i: "   -6.022f23  ", e: false},
		{i: "   -a 6.022e23  ", e: false},
		{i: "   - 6.022.23  ", e: false},
		{i: "   .0 ", e: true},
		{i: "   . ", e: false},
		{i: "   1. ", e: true},
		{i: "   1.0 ", e: true},
	}
	for _, t := range ts {
		var a = isNumber(t.i)
		if t.e != a {
			fmt.Printf("isNumber(%v) should have been %v but was %v\n", t.i, t.e, a)
		} else {
			fmt.Printf("This test case: %v was OK!\n", t.i)
		}
	}
}
