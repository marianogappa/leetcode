package main

import "fmt"

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	var (
		l = map[byte]byte{
			'A': 'a',
			'B': 'b',
			'C': 'c',
			'D': 'd',
			'E': 'e',
			'F': 'f',
			'G': 'g',
			'H': 'h',
			'I': 'i',
			'J': 'j',
			'K': 'k',
			'L': 'l',
			'M': 'm',
			'N': 'n',
			'O': 'o',
			'P': 'p',
			'Q': 'q',
			'R': 'r',
			'S': 's',
			'T': 't',
			'U': 'u',
			'V': 'v',
			'W': 'w',
			'X': 'x',
			'Y': 'y',
			'Z': 'z',
			'a': 'a',
			'b': 'b',
			'c': 'c',
			'd': 'd',
			'e': 'e',
			'f': 'f',
			'g': 'g',
			'h': 'h',
			'i': 'i',
			'j': 'j',
			'k': 'k',
			'l': 'l',
			'm': 'm',
			'n': 'n',
			'o': 'o',
			'p': 'p',
			'q': 'q',
			'r': 'r',
			's': 's',
			't': 't',
			'u': 'u',
			'v': 'v',
			'w': 'w',
			'x': 'x',
			'y': 'y',
			'z': 'z',
			'0': '0',
			'1': '1',
			'2': '2',
			'3': '3',
			'4': '4',
			'5': '5',
			'6': '6',
			'7': '7',
			'8': '8',
			'9': '9',
		}
		i      = 0
		j      = len(s) - 1
		bi, bj byte
	)
	for i < j {
		bi, i = nextFromLeft(s, i, l)
		bj, j = nextFromRight(s, j, l)
		if i >= j {
			return true
		}
		if bi != bj {
			return false
		}
		i++
		j--
	}
	return true
}

func nextFromLeft(s string, i int, l map[byte]byte) (byte, int) {
	for i < len(s)-1 {
		if li, ok := l[s[i]]; ok {
			return li, i
		}
		i++
	}
	return s[len(s)-1], len(s) - 1
}

func nextFromRight(s string, j int, l map[byte]byte) (byte, int) {
	for j >= 0 {
		if lj, ok := l[s[j]]; ok {
			return lj, j
		}
		j--
	}
	return s[0], 0
}

func main() {
	var ts = []struct {
		i string
		e bool
	}{
		{i: "A man, a plan, a canal: Panama", e: true},
		{i: "race a car", e: false},
		{i: "", e: true},
		{i: ".", e: true},
		{i: "A", e: true},
		{i: "Aa", e: true},
		{i: "aA", e: true},
		{i: "aBA", e: true},
		{i: "aBbA", e: true},
		{i: "aBcbA", e: true},
		{i: "aBcdbA", e: false},
		{i: "0P", e: false},
	}
	for _, t := range ts {
		var a = isPalindrome(t.i)
		if t.e != a {
			fmt.Printf("isPalindrome(%v) should have been %v but was ", t.i, t.e, a)
		}
	}
}
