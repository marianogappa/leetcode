package main

import "fmt"

func addBinary(a string, b string) string {
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}
	var (
		r     = make([]byte, 0, max(len(a), len(b))+1)
		carry = 0
		c     = ""
		byt   byte
	)
	for i := 0; i < min(len(a), len(b)); i++ {
		byt, carry = alu(a[len(a)-1-i], b[len(b)-1-i], carry)
		r = append(r, byt)
	}
	if len(a) > len(b) {
		c = a[:len(a)-len(b)]
	}
	if len(b) > len(a) {
		c = b[:len(b)-len(a)]
	}
	for i := len(c) - 1; i >= 0; i-- {
		byt, carry = alu(c[i], '0', carry)
		r = append(r, byt)
	}
	if carry == 1 {
		r = append(r, '1')
	}
	return reverse(string(r))
}

func alu(a, b byte, carry int) (byte, int) {
	var r byte
	count := carry
	if a == '1' {
		count++
	}
	if b == '1' {
		count++
	}
	if count == 1 || count == 3 {
		r = '1'
	} else {
		r = '0'
	}
	carry = 0
	if count > 1 {
		carry = 1
	}
	return r, carry
}

func reverse(r string) string {
	if r == "" {
		return ""
	}
	var (
		t = []byte(r)
		i = 0
		j = len(r) - 1
	)
	for i < j {
		t[i], t[j] = t[j], t[i]
		i++
		j--
	}
	return string(t)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ts := []struct {
		a        string
		b        string
		expected string
	}{
		{"", "", ""},
		{"", "0", "0"},
		{"1", "", "1"},
		{"00", "00", "00"},
		{"01", "00", "01"},
		{"00", "01", "01"},
		{"10", "00", "10"},
		{"00", "10", "10"},
		{"10", "01", "11"},
		{"11", "01", "100"},
		{"11", "11", "110"},
		{"11", "100", "111"},
		{"11", "111", "1010"},
		{"11", "10000", "10011"},
	}
	for _, t := range ts {
		if addBinary(t.a, t.b) != t.expected {
			fmt.Printf("addBinary(%v, %v) should have been %v but was %v\n", t.a, t.b, t.expected, addBinary(t.a, t.b))
		}
	}
}
