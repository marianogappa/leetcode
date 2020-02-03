package main

import "fmt"

func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	var (
		bs      = make(map[byte]int, 0)
		counter int
	)

	// Map wanted frequencies
	for i := 0; i < len(t); i++ {
		bs[t[i]]++
		counter++
	}
	var (
		start, bestStart, end int
		bestEnd               = len(s)
		found                 bool
	)
	for i := 0; i < len(s); i++ {
		bs[s[i]]-- // If it's not wanted it will be negative
		if bs[s[i]] >= 0 {
			counter--
		}
		end++
		for counter == 0 {
			found = true
			if end-start < bestEnd-bestStart {
				bestStart, bestEnd = start, end
			}
			start++
			bs[s[start-1]]++
			if bs[s[start-1]] > 0 {
				counter++
				break
			}
		}
	}
	if !found {
		return ""
	}
	return s[bestStart:bestEnd]
}

func main() {
	var ts = []struct {
		s, t string
		e    string
	}{
		{s: "", t: "", e: ""},
		{s: "ADOBECODEBANC", t: "ABC", e: "BANC"},
		{s: "a", t: "aa", e: ""},
		{s: "aaabcbda", t: "abc", e: "abc"},
		{s: "aaabdcbca", t: "abc", e: "bca"},
		{s: "cabwefgewcwaefgcf", t: "cae", e: "cwae"},
		{s: "a", t: "a", e: "a"},
	}
	for _, t := range ts {
		var a = minWindow(t.s, t.t)
		if t.e != a {
			fmt.Printf("minWindow(%v, %v) should have been %v but was %v\n", t.s, t.t, t.e, a)
		}
	}
}
