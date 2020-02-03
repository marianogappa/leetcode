package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var ls = make([]int, 26)
	for i := 0; i < len(s); i++ {
		ls[s[i]-'a']++
		ls[t[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if ls[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("", "") == true)
	fmt.Println(isAnagram("a", "a") == true)
	fmt.Println(isAnagram("anagram", "nagaram") == true)
	fmt.Println(isAnagram("rat", "car") == false)
}
