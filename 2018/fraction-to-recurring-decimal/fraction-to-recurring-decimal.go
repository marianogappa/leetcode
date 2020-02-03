package main

import "fmt"

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	var r = make([]byte, 0)
	if numerator*denominator < 0 {
		r = append(r, '-')
	}
	var (
		num = int64(abs(numerator))
		den = int64(abs(denominator))
	)
	var integerPart = num / den
	r = append(r, []byte(fmt.Sprintf("%v", integerPart))...)
	num -= integerPart * den
	if num == 0 {
		return string(r)
	}
	r = append(r, '.')

	var repeated = make(map[int64]int, 0)
	repeated[num] = len(r)
	for num > 0 {
		num *= 10
		r = append(r, []byte(fmt.Sprintf("%v", num/den))...)
		num %= den

		if n, ok := repeated[num]; ok {
			var per = make([]byte, 0)
			per = append(per, r[n:]...)
			r = r[:n]
			r = append(r, []byte(fmt.Sprintf("(%v)", string(per)))...)
			break
		} else {
			repeated[num] = len(r)
		}
	}
	return string(r)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var fs = [][]int{
		{1, 2},
		{2, 4},
		{4, 2},
		{0, 1},
		{0, -1},
		{1, -2},
		{-1, 2},
		{-1, -2},
		{3, 8},
		{1, 3},
		{1, 30},
		{1, 333},
		{1, 9801},
		{1, 90},
		{1, 214748364},
	}
	for _, f := range fs {
		fmt.Printf("fractionToDecimal(%v/%v) = %v\n", f[0], f[1], fractionToDecimal(f[0], f[1]))
	}
}
