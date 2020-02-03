package main

import "fmt"

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	var i = 0
	for i < len(str) && (str[i] == ' ' || str[i] == '\t' || str[i] == '\n' || str[i] == '\r') {
		i++
	}
	if i == len(str) {
		return 0
	}
	var sign = 1
	if str[i] == '+' {
		i++
	} else if str[i] == '-' {
		sign = -1
		i++
	}
	var nums = make([]byte, 0)
	for i < len(str) && str[i] >= '0' && str[i] <= '9' {
		nums = append(nums, str[i])
		i++
	}
	if len(nums) == 0 {
		return 0
	}
	if len(nums) > 10 || (len(nums) == 10 && ((sign == 1 && isGreater(nums, []byte("2147483647"))) || (sign == -1 && isGreater(nums, []byte("2147483648"))))) {
		if sign == -1 {
			return -2147483648
		}
		return 2147483647
	}
	var (
		n = 0
		d = 1
	)
	for i := len(nums) - 1; i >= 0; i-- {
		n += int(nums[i]-'0') * d
		d *= 10
	}
	return n * sign
}

func isGreater(a, b []byte) bool {
	if len(a) == 0 {
		return false
	}
	return a[0] > b[0] || (a[0] == b[0] && isGreater(a[1:], b[1:]))
}

func main() {
	var ts = []struct {
		i string
		e int
	}{
		{i: "", e: 0},
		{i: " ", e: 0},
		{i: "  ", e: 0},
		{i: "  \n", e: 0},
		{i: "  \n\r", e: 0},
		{i: "  \n\r\t", e: 0},
		{i: "  \n\r\t1", e: 1},
		{i: "  \n\r\t+1", e: 1},
		{i: "  \n\r\t-1", e: -1},
		{i: "  \n\r\t+-1", e: 0},
		{i: " a \n\r\t+1", e: 0},
		{i: "   \n\r\t+1aaaaa", e: 1},
		{i: "   \n\r\t-1aaaaa", e: -1},
		{i: "   \n\r\t-1.aaaaa", e: -1},
		{i: "12345678901", e: 2147483647},
		{i: "-12345678901", e: -2147483648},
		{i: "2147483646", e: 2147483646},
		{i: "-2147483647", e: -2147483647},
		{i: "1095502006p8", e: 1095502006},
	}
	for _, t := range ts {
		var a = myAtoi(t.i)
		if t.e != a {
			fmt.Printf("myAtoi(%v) should have been %v but was %v\n", t.i, t.e, a)
		}
	}
}
