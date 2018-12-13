package main

import "fmt"

func isValid(s string) bool {
    var closers = map[rune]rune{'(': ')', '{': '}', '[': ']'}
    var opened = make([]rune, 0)
    for _, r := range s {
        if r == ')' || r == '}' || r == ']' {
            if len(opened) == 0 {
                return false
            }
            if c, ok := closers[opened[len(opened)-1]]; !ok || c != r {
                return false
            }
            opened = opened[:len(opened)-1]
        } else {
            opened = append(opened, r)
        }
    }
    if len(opened) != 0 {
        return false
    }
    return true
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
}
