package main

import "fmt"

func licenseKeyFormatting(S string, K int) string {
	if len(S) == 0 {
		return ""
	}
	var (
		c int
		r = make([]byte, 0, len(S))
	)
	for i := 0; i < len(S); i++ {
		if S[i] != '-' {
			c++
		}
	}
	var i = 0
	for j := 0; j < c%K; j++ {
		for S[i] == '-' {
			i++
		}
		r = append(r, upper(S[i]))
		i++
	}
	for j := 0; j < c/K; j++ {
		if j > 0 || c%K > 0 {
			r = append(r, '-')
		}
		for k := 0; k < K; k++ {
			for S[i] == '-' {
				i++
			}
			r = append(r, upper(S[i]))
			i++
		}
	}
	return string(r)
}

func upper(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b - 'a' + 'A'
	}
	return b
}

func main() {
	var ts = []struct {
		S string
		K int
		e string
	}{
		{S: "5F3Z-2e-9-w", K: 4, e: "5F3Z-2E9W"},
		{S: "5F3Z2e-9w", K: 4, e: "5F3Z-2E9W"},
		{S: "2-5g-3-J", K: 2, e: "2-5G-3J"},
		{S: "2-5g-3-J", K: 1, e: "2-5-G-3-J"},
		{S: "", K: 2, e: ""},
		{S: "a", K: 2, e: "A"},
	}
	for _, t := range ts {
		var a = licenseKeyFormatting(t.S, t.K)
		if t.e != a {
			fmt.Printf("licenseKeyFormatting(%v, %v) should have been %v but was %v\n", t.S, t.K, t.e, a)
		}
	}
}
