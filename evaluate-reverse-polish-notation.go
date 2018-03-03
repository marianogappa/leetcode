package main

import "fmt"

func readNum(s string) int {
	if len(s) == 0 {
		return 0
	}
	var (
		n    = 0
		d    = 1
		sign = 1
	)
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	}
	for i := len(s) - 1; i >= 0; i-- {
		n += int(s[i]-'0') * d
		d *= 10
	}
	return n * sign
}

func evalRPN(tokens []string) int {
	var (
		ns = make([]int, 0)
		n  int
	)
	for _, t := range tokens {
		switch t {
		case "+":
			n = ns[len(ns)-2] + ns[len(ns)-1]
		case "*":
			n = ns[len(ns)-2] * ns[len(ns)-1]
		case "-":
			n = ns[len(ns)-2] - ns[len(ns)-1]
		case "/":
			n = ns[len(ns)-2] / ns[len(ns)-1]
		default:
			ns = append(ns, readNum(t))
			continue
		}
		ns = ns[:len(ns)-2]
		ns = append(ns, n)
	}
	return ns[len(ns)-1]
}

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}) == 9)
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}) == 6)
	fmt.Println(evalRPN([]string{"-1"}) == -1)
}
