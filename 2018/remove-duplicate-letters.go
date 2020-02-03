package main

import "fmt"

func removeDuplicateLetters(s string) string {
	if len(s) <= 1 {
		return s
	}
	var fs = make(map[byte]int, 0) // find letter frequencies
	for i := 0; i < len(s); i++ {
		fs[s[i]]++
	}
	var smallest = 0 // find smallest
	for i := 0; i < len(s); i++ {
		if s[i] < s[smallest] {
			smallest = i
		}
		fs[s[i]]--
		if fs[s[i]] == 0 { // don't choose a smaller after a unique
			break
		}
	}
	var bs = make([]byte, 0) // trim rest of the string of the chosen smallest
	for i := smallest + 1; i < len(s); i++ {
		if s[i] == s[smallest] {
			continue
		}
		bs = append(bs, s[i])
	}
	return string(s[smallest]) + removeDuplicateLetters(string(bs)) // result is smallest + solving the rest
}

func main() {
	var ts = []struct {
		i string
		e string
	}{
		{i: "", e: ""},
		{i: "a", e: "a"},
		{i: "abc", e: "abc"},
		{i: "abac", e: "abc"},
		{i: "ababc", e: "abc"},
		{i: "aaaaaa", e: "a"},
		{i: "hablaba", e: "habl"},
		{i: "hbalaba", e: "halb"},
		{i: "hbaclaba", e: "haclb"},
		{i: "bcabc", e: "abc"},
		{i: "cbacdcbc", e: "acdb"},
		{i: "abacb", e: "abc"},
		{i: "bccab", e: "bca"},
	}
	for _, t := range ts {
		var a = removeDuplicateLetters(t.i)
		if t.e != a {
			fmt.Printf("removeDuplicateLetters(%v) should have been %v but was %v\n", t.i, t.e, a)
		}
	}
}
