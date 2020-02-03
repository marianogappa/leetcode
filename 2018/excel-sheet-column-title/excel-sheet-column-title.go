package main

import (
	"fmt"
)

func convertToTitle(n int) string {
	if n == 0 {
		return ""
	}
	return convertToTitle((n-1)/26) + string('A'+(n-1)%26)
}

func main() {
	var ts = []struct {
		n int
		e string
	}{
		{n: 1, e: "A"},
		{n: 2, e: "B"},
		{n: 26, e: "Z"},
		{n: 27, e: "AA"},
		{n: 28, e: "AB"},
		{n: 676, e: "ZA"},
		{n: 677, e: "ZB"},
	}
	for _, t := range ts {
		var a = convertToTitle(t.n)
		if t.e != a {
			fmt.Printf("convertToTitle(%v) should have been %v but was %v\n", t.n, t.e, a)
		}
	}
}
