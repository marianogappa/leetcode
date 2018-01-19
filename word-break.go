package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	return doWB(0, make(map[int]bool, 0), s, wordDict)
}

func doWB(i int, c map[int]bool, s string, ws []string) bool {
	if i == len(s) {
		return true
	}
	if r, ok := c[i]; ok {
		return r
	}
	var res = false
	for _, w := range ws {
		if isPrefix(w, s[i:]) && doWB(i+len(w), c, s, ws) {
			res = true
		}
	}
	c[i] = res
	return res
}

func isPrefix(sb, s string) bool {
	if len(sb) > len(s) {
		return false
	}
	for i := range sb {
		if sb[i] != s[i] {
			return false
		}
	}
	return true
}

func main() {
	var ts = []struct {
		s        string
		wordDict []string
		e        bool
	}{
		{
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			e:        true,
		},
	}
	for _, t := range ts {
		var a = wordBreak(t.s, t.wordDict)
		if t.e != a {
			fmt.Printf("wordBreak(%v, %v) should have been %v but was %v\n", t.s, t.wordDict, t.e, a)
		}
	}
}
