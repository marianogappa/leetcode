package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	var m, r = make(map[byte]byte, 0), make(map[byte]byte, 0)
	for i := 0; i < len(s); i++ {
		if b, ok := m[s[i]]; ok && b != t[i] {
			return false // if entry for s[i] already exists and maps to different
		}
		if rb, ok := r[t[i]]; ok && rb != s[i] {
			return false // if mapping for s[i] is used for a different
		}
		m[s[i]] = t[i]
		r[t[i]] = s[i]
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add") == true)
	fmt.Println(isIsomorphic("foo", "bar") == false)
	fmt.Println(isIsomorphic("paper", "title") == true)
	fmt.Println(isIsomorphic("ab", "aa") == false)
}
