package main

import (
	"fmt"
	"reflect"
	"sort"
)

func wordBreak(s string, wordDict []string) []string {
	return doWB(0, make(map[int][]string, 0), s, wordDict)
}

func doWB(i int, c map[int][]string, s string, ws []string) []string {
	if i >= len(s) {
		return []string{""}
	}
	if r, ok := c[i]; ok {
		return r
	}
	var res = make([]string, 0)
	for _, w := range ws {
		if isPrefix(w, s[i:]) {
			var rcs = doWB(i+len(w), c, s, ws)
			for _, rc := range rcs {
				var space = " "
				if rc == "" {
					space = ""
				}
				res = append(res, w+space+rc)
			}
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
		e        []string
	}{
		{
			s:        "catsanddog",
			wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			e:        []string{"cats and dog", "cat sand dog"},
		},
		{
			s:        "",
			wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			e:        []string{""},
		},
	}
	for _, t := range ts {
		var a = wordBreak(t.s, t.wordDict)
		sort.Strings(a)
		sort.Strings(t.e)
		if !reflect.DeepEqual(t.e, a) {
			fmt.Printf("wordBreak(%v, %v) should have been %v but was %v\n", t.s, t.wordDict, t.e, a)
		}
	}
}
