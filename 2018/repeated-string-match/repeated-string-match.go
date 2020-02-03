package main

import "fmt"

func repeatedStringMatch(A string, B string) int {
	if B == "" {
		return 0
	}
	if A == "" {
		return -1
	}
	var i, j int
	for {
		if j == len(B) {
			var extra = 0
			if i%len(A) > 0 {
				extra++
			}
			return i/len(A) + extra
		}
		if A[i%len(A)] == B[j] {
			j++
		} else {
			j = 0
			if i >= len(A) {
				return -1
			}
		}
		i++
	}
}

func main() {
	var ts = []struct {
		A, B string
		e    int
	}{
		{A: "", B: "", e: 0},
		{A: "", B: "a", e: -1},
		{A: "abcd", B: "cdabcdab", e: 3},
		{A: "baaaaaaaaaaaaa", B: "ab", e: 2},
	}
	for _, t := range ts {
		var a = repeatedStringMatch(t.A, t.B)
		if t.e != a {
			fmt.Printf("repeatedStringMatch(%v, %v) should have been %v but was %v\n", t.A, t.B, t.e, a)
		}
	}
}
