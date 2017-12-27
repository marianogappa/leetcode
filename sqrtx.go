package main

import "fmt"

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	var n int
	for n = 2; n*n <= x; n++ {
	}
	return n - 1
}

func main() {
	var ts = []struct {
		i, e int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 2},
		{6, 2},
		{7, 2},
		{8, 2},
		{9, 3},
		{123456789, 11111},
		{1234567890, 35136},
	}
	for _, t := range ts {
		var a = mySqrt(t.i)
		if t.e != a {
			fmt.Printf("mySqrt(%v) should have been %v but was %v\n", t.i, t.e, a)
		}
	}
}
