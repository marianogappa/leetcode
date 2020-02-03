package main

import "fmt"

func isMatch(s string, p string) bool {
	fmt.Printf("isMatch(%v, %v)\n", s, p)
	var pa *pat
	var newP string
	pa, newP = nextPat(p)
	if pa == nil {
		return s == ""
	}
	if s == "" {
		return pa.star && (newP == "" || isMatch(s, newP))
	}
	if pa.dot || s[0] == pa.lit {
		if pa.star {
			return isMatch(s[1:], p) || isMatch(s, newP)
		} else {
			return isMatch(s[1:], newP)
		}
	} else if pa.star {
		return isMatch(s, newP)
	} else {
		return false
	}
}

type pat struct {
	dot, star bool
	lit       byte
}

func nextPat(p string) (*pat, string) {
	if p == "" {
		return nil, ""
	}
	pa := pat{lit: p[0]}
	if p[0] == '.' {
		pa.dot = true
	}
	if len(p) > 1 && p[1] == '*' {
		pa.star = true
		return &pa, p[2:]
	}
	return &pa, p[1:]
}

/*
isMatch("aa","a") → false
isMatch("aa","aa") → true
isMatch("aaa","aa") → false
isMatch("aa", "a*") → true
isMatch("aa", ".*") → true
isMatch("ab", ".*") → true
isMatch("aab", "c*a*b") → true
*/
func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "aa"))
	fmt.Println(isMatch("aaa", "aa"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("aa", ".*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("aaa", "a*a"))
	fmt.Println(isMatch("", "c*c*"))
}
