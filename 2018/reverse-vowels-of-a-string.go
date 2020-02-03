package main

import "fmt"

func firstVowel(s string, i int) int {
	for j := i; j < len(s); j++ {
		switch s[j] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			return j
		}
	}
	return -1
}

func lastVowel(s string, i int) int {
	for j := i; j >= 0; j-- {
		switch s[j] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			return j
		}
	}
	return -1
}

func reverseVowels(s string) string {
	var (
		f = 0
		l = len(s) - 1
		b = []byte(s)
	)
	for {
		f = firstVowel(s, f)
		l = lastVowel(s, l)
		if f == -1 || f > l {
			break
		}
		b[f], b[l] = b[l], b[f]
		f++
		l--
	}
	return string(b)
}

func main() {
	var ts = []struct {
		i, expected string
	}{
		{"", ""},
		{"ae", "ea"},
		{"aA", "Aa"},
		{"ae", "ea"},
		{"hello", "holle"},
		{"leetcode", "leotcede"},
		{"bcd", "bcd"},
		{"aa", "aa"},
		{"aae", "eaa"},
	}
	for _, t := range ts {
		var a = reverseVowels(t.i)
		if a != t.expected {
			fmt.Printf("reverseVowels(%v) should have been %v but was %v\n", t.i, t.expected, a)
		}
	}
}
