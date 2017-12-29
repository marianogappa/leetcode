package main

import "fmt"

func areLinked(a, b string) bool {
	var c int
	for i := 0; i < len(a) && c <= 1; i++ {
		if a[i] != b[i] {
			c++
		}
	}
	return c == 1
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var prereq bool
	for _, w := range wordList {
		if w == endWord {
			prereq = true
			break
		}
	}
	if !prereq {
		return 0
	}

	if areLinked(beginWord, endWord) {
		return 2
	}

	var linksToEndWord = map[string]int{endWord: 2}
	traverse(wordList, linksToEndWord, 3)

	var min = 1<<31 - 1
	for w, cd := range linksToEndWord {
		if areLinked(beginWord, w) && cd < min {
			min = cd
		}
	}
	if min < 1<<31-1 {
		return min
	}
	return 0
}

func traverse(wordList []string, ls map[string]int, d int) {
	var ws = make([]string, 0)
	for cw, c := range ls {
		if c == d-1 {
			for i, w := range wordList {
				if w != "" && w != cw && areLinked(cw, w) {
					ws = append(ws, w)
					wordList[i] = ""
				}
			}
		}
	}
	for _, w := range ws {
		ls[w] = d
	}
	if len(ws) > 0 {
		traverse(wordList, ls, d+1)
	}
}

func main() {
	var ts = []struct {
		b, e string
		ws   []string
		ex   int
	}{
		{b: "hit", e: "cog", ws: []string{"hot", "dot", "dog", "lot", "log", "cog"}, ex: 5},
		{b: "hit", e: "zzz", ws: []string{"hot", "dot", "dog", "lot", "log", "cog", "zzz"}, ex: 0},
		{b: "hit", e: "cog", ws: []string{"hot", "hoa", "xoa", "xob", "xoc", "xod", "xoe", "xof", "xog", "dot", "dog", "lot", "log", "cog"}, ex: 5},
		{b: "hit", e: "his", ws: []string{"his"}, ex: 2},
		{b: "hit", e: "hfs", ws: []string{"his", "hfs"}, ex: 3},
		{b: "hit", e: "cog", ws: []string{"hot", "dot", "dog", "lot", "log"}, ex: 0},
		{b: "qa", e: "sq", ws: []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}, ex: 5},
	}
	for i, t := range ts {
		var a = ladderLength(t.b, t.e, t.ws)
		if t.ex != a {
			fmt.Printf("ladderLength(%v, %v, %v) should have been %v but was %v\n", t.b, t.e, t.ws, t.ex, a)
		} else {
			fmt.Println("Test case", i, "was ok")
		}
	}
}
