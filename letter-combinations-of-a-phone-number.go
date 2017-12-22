package main

import (
	"fmt"
	"reflect"
	"sort"
)

var m = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	return doLetterCombinations(digits, make(map[string][]string, 0))
}

func doLetterCombinations(digits string, c map[string][]string) []string {
	if digits == "" {
		return []string{""}
	}
	if _, ok := m[digits[0]]; !ok {
		return doLetterCombinations(digits[1:], c)
	}
	if cr, ok := c[digits]; ok {
		return cr
	}
	var cs = doLetterCombinations(digits[1:], c)
	var newCs = make([]string, len(cs)*len(m[digits[0]]))
	i := 0
	for _, p := range cs {
		for _, d := range m[digits[0]] {
			newCs[i] = string(d) + p
			i++
		}
	}
	c[digits] = newCs
	return newCs
}

func main() {
	var ts = []struct {
		i string
		e []string
	}{
		{i: "", e: []string{}},
		{i: "23", e: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{i: "3", e: []string{"d", "e", "f"}},
	}
	for _, t := range ts {
		var a = letterCombinations(t.i)
		sort.Strings(a)
		sort.Strings(t.e)
		if !reflect.DeepEqual(t.e, a) {
			fmt.Printf("letterCombinations(%v) should have been %v but was %v\n", t.i, t.e, a)
		}
	}
}
