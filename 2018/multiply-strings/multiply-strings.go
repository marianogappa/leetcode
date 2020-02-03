package main

import "fmt"

func getNum(b byte) int {
	switch b {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	default:
		return 9
	}
}

func getNums(n string) []int {
	var ns = make([]int, len(n))
	for i := 0; i < len(n); i++ {
		ns[i] = getNum(n[len(n)-1-i])
	}
	return ns
}

func reverse(b []byte) {
	var (
		i = 0
		j = len(b) - 1
	)
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
}

func multiply(num1 string, num2 string) string {
	var (
		ns1, ns2 = getNums(num1), getNums(num2)
		rs       = make(map[int]int)
		r        = make([]byte, 0)
	)
	for i1, n1 := range ns1 {
		for i2, n2 := range ns2 {
			rs[i1+i2] += n1 * n2
		}
	}
	var carry = 0
	for i := 0; i <= len(ns1)+len(ns2)-2; i++ {
		rs[i] += carry
		var m = rs[i] % 10
		r = append(r, []byte(fmt.Sprintf("%v", m))[0])
		carry = rs[i] - m
		carry /= 10
	}
	if carry > 0 { // add trailing carry
		var carryB = []byte(fmt.Sprintf("%v", carry))
		reverse(carryB)
		r = append(r, carryB...)
	}
	reverse(r)
	var i int // reslice to remove leading zeroes
	for i = 0; i < len(r)-1; i++ {
		if r[i] != '0' {
			break
		}
	}
	return string(r[i:])
}

func main() {
	var ts = []struct {
		num1, num2, expected string
	}{
		{num1: "2", num2: "3", expected: "6"},
		{num1: "20", num2: "3", expected: "60"},
		{num1: "20", num2: "300", expected: "6000"},
		{num1: "0", num2: "300", expected: "0"},
		{num1: "010", num2: "200", expected: "2000"},
		{num1: "9", num2: "9", expected: "81"},
		{num1: "99", num2: "99", expected: "9801"},
	}
	for _, t := range ts {
		var actual = multiply(t.num1, t.num2)
		if t.expected != actual {
			fmt.Printf("multiply(%v, %v) should have been %v but was %v\n", t.num1, t.num2, t.expected, actual)
		}
	}
}
