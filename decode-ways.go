package main

import "fmt"

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) > 0 && s[0] == '0' {
		return 0
	}
	var c = map[int]int{len(s): 1}
	return doDecodings(s, 0, c)
}

func doDecodings(s string, i int, c map[int]int) int {
	if n, ok := c[i]; ok {
		return n
	}
	if s[i] == '0' {
		return 0
	}
	var r int
	if i+1 < len(s) && (s[i] == '1' || (s[i] == '2' && s[i+1] >= '0' && s[i+1] <= '6')) {
		if s[i+1] == '0' {
			r = doDecodings(s, i+2, c)
		} else {
			r = doDecodings(s, i+1, c) + doDecodings(s, i+2, c)
		}
		c[i] = r
		return r
	}
	r = doDecodings(s, i+1, c)
	c[i] = r
	return r
}

func main() {
	var ts = []struct {
		s string
		e int
	}{
		{s: "", e: 0},
		{s: "1", e: 1},
		{s: "3", e: 1},
		{s: "11", e: 2},
		{s: "111", e: 3},
		{s: "121", e: 3},
		{s: "291", e: 1},
		{s: "261", e: 2},
		{s: "126", e: 3},
		{s: "106", e: 1},
		{s: "610", e: 1},
		{s: "2610", e: 2},
		{s: "0", e: 0},
		{s: "00", e: 0},
		{s: "000000", e: 0},
		{s: "0000001", e: 0},
		{s: "0000002610", e: 0},
		{s: "11111111111111111111111111111111111111111111111111111111111111111111111111111", e: 8944394323791464},
	}
	for _, t := range ts {
		var a = numDecodings(t.s)
		if t.e != a {
			fmt.Printf("numDecodings(%v) should have been %v but was %v\n", t.s, t.e, a)
		}
	}
}
