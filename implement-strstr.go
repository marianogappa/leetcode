package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	var (
		j  int
		j0 = -1
	)
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[j] {
			if j0 == -1 {
				j0 = i
			}
			j++
			if j == len(needle) {
				return j0
			}
		} else {
			if j > 0 {
				i = j0
			}
			j = 0
			j0 = -1
		}
	}
	return -1
}

func main() {
	ts := []struct {
		h        string
		n        string
		expected int
	}{
		{"hello", "", 0},
		{"", "", 0},
		{"", "h", -1},
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"aaaaaaaaaaaazaaa", "z", 12},
		{"aaa", "aaaa", -1},
		{"mississippi", "issip", 4},
		{"aaazaaaaz", "aaaaz", 4},
		{"mississippi", "pi", 9},
		{"missipissip", "pi", 5},
	}

	for _, t := range ts {
		if strStr(t.h, t.n) != t.expected {
			fmt.Printf("strStr(%v, %v) should have been %v but was %v\n", t.h, t.n, t.expected, strStr(t.h, t.n))
		}
	}
}
