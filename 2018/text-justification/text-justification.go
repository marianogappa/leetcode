package main

import (
	"fmt"
	"reflect"
)

func fullJustify(words []string, maxWidth int) []string {
	if len(words) == 0 { // special case; empty
		return []string{}
	}
	var (
		start = 0
		wsl   = 0
		res   = make([]string, 0)
	)
	for i, w := range words {
		var wl = len(w)
		if start < i {
			wl++
		}
		if wsl+wl > maxWidth { // If exceeded line length
			res = append(res, justifyLine(words[start:i], wsl, maxWidth))
			start = i
			wsl = len(w)
			continue
		}
		wsl += wl
	}
	var rest = make([]byte, 0)
	for i := start; i < len(words); i++ {
		rest = append(rest, []byte(words[i])...)
		if i < len(words)-1 {
			rest = append(rest, ' ')
		}
	}
	for len(rest) < maxWidth {
		rest = append(rest, ' ')
	}
	res = append(res, string(rest))
	return res
}

func justifyLine(ws []string, wsl, ll int) string {
	// Special case when just one word (because / 0)
	if len(ws) == 1 {
		var res = []byte(ws[0])
		for len(res) < ll {
			res = append(res, ' ')
		}
		return string(res)
	}
	var (
		extra         = ll - wsl
		maxAll        = extra / (len(ws) - 1) // spaces are len(ws)-1!
		extraFromLeft = extra % (len(ws) - 1)
		res           = make([]byte, 0, ll)
	)
	for i, w := range ws {
		res = append(res, w...)
		if i < len(ws)-1 { // except last word, put mandatory space
			for i := 0; i < maxAll; i++ {
				res = append(res, ' ')
			}
			res = append(res, ' ')
		}
		if extraFromLeft > 0 { // extra uneven goes to left
			res = append(res, ' ')
			extraFromLeft--
		}
	}
	return string(res)
}

func main() {
	var ts = []struct {
		words    []string
		maxWidth int
		expected []string
	}{
		{
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
			expected: []string{
				"This    is    an",
				"example  of text",
				"justification.  "},
		},
		{
			words:    []string{"Only"},
			maxWidth: 5,
			expected: []string{
				"Only "},
		},
		{
			words:    []string{"Only"},
			maxWidth: 4,
			expected: []string{
				"Only"},
		},
		{
			words:    []string{"Only", "few", "lines"},
			maxWidth: 5,
			expected: []string{
				"Only ", "few  ", "lines"},
		},
		{
			words:    []string{"Only", "few", "lines"},
			maxWidth: 10,
			expected: []string{
				"Only   few", "lines     "},
		},
		{
			words:    []string{},
			maxWidth: 10,
			expected: []string{},
		},
	}
	for _, t := range ts {
		var a = fullJustify(t.words, t.maxWidth)
		if !reflect.DeepEqual(t.expected, a) {
			fmt.Printf("fullJustify(%v, %v) should have been %v but was %v\n", t.words, t.maxWidth, t.expected, a)
		}
	}
}
