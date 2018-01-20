package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	return bfs(&s, len(s))
}

func bfs(s *string, l int) string {
	if l == 1 {
		return string((*s)[0])
	}
	for i := 0; i+l <= len(*s); i++ {
		if isPalindrome(s, i, i+l-1) {
			return (*s)[i : i+l]
		}
	}
	return bfs(s, l-1)
}

func isPalindrome(s *string, i, j int) bool {
	for i < j {
		if (*s)[i] != (*s)[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(longestPalindrome("babad") == "bab")
	fmt.Println(longestPalindrome("cbbd") == "bb")
	fmt.Println(longestPalindrome("a") == "a")
	fmt.Println(longestPalindrome("bb") == "bb")
	fmt.Println(longestPalindrome("abcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabc") == "a")
}
