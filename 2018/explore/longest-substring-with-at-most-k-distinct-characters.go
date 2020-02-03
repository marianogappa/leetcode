package main

import (
	"fmt"
)

func lengthOfLongestSubstringKDistinct(s string, k int) int {
    if len(s) == 0 || k == 0 {
        return 0
    }
    var (
        bestStart, bestEnd, start int
        letters = make(map[byte]int)
    )
    for end := 0; end < len(s); end++ {
        letters[s[end]]++
        if len(letters) <= k && end-start > bestEnd-bestStart {
            bestStart, bestEnd = start, end
        }
        for start < len(s) && len(letters) > k {
            letters[s[start]]--
            if letters[s[start]] == 0 {
                delete(letters, s[start])
            }
            start++
        }
    }
    return bestEnd-bestStart+1
}

func main() {
    fmt.Println(lengthOfLongestSubstringKDistinct("eceba", 2)==3)
    fmt.Println(lengthOfLongestSubstringKDistinct("aaaaaa", 2)==6)
    fmt.Println(lengthOfLongestSubstringKDistinct("aaaaabcccccc", 2)==7)
}