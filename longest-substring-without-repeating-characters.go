package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	var (
		start, end, maxEnd int
		maxStart           = 1
		fqs                = make(map[byte]struct{}, 0)
	)
	for end = 0; end < len(s); end++ {
		if _, ok := fqs[s[end]]; ok {
			if end-1-start+1 > maxEnd-maxStart+1 {
				maxEnd, maxStart = end-1, start
			}
			for start = start; s[start] != s[end]; start++ {
				delete(fqs, s[start])
			}
			start++
		}
		fqs[s[end]] = struct{}{}
	}
	if end-start+1 > maxEnd-maxStart+1 {
		maxEnd, maxStart = end-1, start
	}
	return maxEnd - maxStart + 1
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
