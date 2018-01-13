package main

import "fmt"

type exp struct {
	ns  []int
	ops []int
}

func (e exp) val() int {
	if len(e.ns) == 0 {
		return 0
	}
	if len(e.ns) == 1 {
		return e.ns[0]
	}
	var r = e.ns[0]
	for i := range e.ops {
		r += e.ns[i+1] * e.ops[i]
	}
	return r
}

func readNum(s string, start int) (int, int) {
	var (
		bs = make([]byte, 0)
		n  = 0
		i  = 0
	)
	for i = start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		bs = append(bs, s[i])
	}
	var d = 1
	for j := len(bs) - 1; j >= 0; j-- {
		n += int(bs[j]-'0') * d
		d *= 10
	}
	return n, i
}

func readParens(s string, pos int) (string, int) {
	var (
		nest = 1
		sub  = make([]byte, 0)
	)
	for nest > 0 {
		pos++
		if s[pos] == '(' {
			nest++
		} else if s[pos] == ')' {
			nest--
		}
		if nest > 0 {
			sub = append(sub, s[pos])
		}
	}
	return string(sub), pos + 1
}

func calculate(s string) int {
	var (
		e        exp
		pos, num int
		sub      string
	)
	for pos < len(s) {
		if s[pos] == '(' {
			sub, pos = readParens(s, pos)
			e.ns = append(e.ns, calculate(sub))
		} else if s[pos] == '+' {
			e.ops = append(e.ops, 1)
			pos++
		} else if s[pos] == '-' {
			e.ops = append(e.ops, -1)
			pos++
		} else if s[pos] >= '0' && s[pos] <= '9' {
			num, pos = readNum(s, pos)
			e.ns = append(e.ns, num)
		} else {
			pos++
		}
	}
	return e.val()
}

func main() {
	fmt.Println(calculate("( 1 + 2) -3") == 0)
	fmt.Println(calculate("") == 0)
	fmt.Println(calculate("9") == 9)
	fmt.Println(calculate("10") == 10)
	fmt.Println(calculate("1234") == 1234)
	fmt.Println(calculate("2-1") == 1)
	fmt.Println(calculate("2     +      1    -      2") == 1)
	fmt.Println(calculate("(((1)))") == 1)
	fmt.Println(calculate("(((1)))") == 1)
	fmt.Println(calculate("((1+(4-2))+2)") == 5)
	fmt.Println(calculate("()") == 0)
}
