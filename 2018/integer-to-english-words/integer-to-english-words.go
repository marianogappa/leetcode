package main

import (
	"fmt"
	"math"
	"strings"
)

var ds = map[int]string{
	30: "Nonillion",
	27: "Octillion",
	24: "Septillion",
	21: "Sixtillion",
	18: "Quintillion",
	15: "Quadrillion",
	12: "Trillion",
	9:  "Billion",
	6:  "Million",
	3:  "Thousand",
}
var to19 = map[int]string{
	0:  "Zero",
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
	11: "Eleven",
	12: "Twelve",
	13: "Thirteen",
	14: "Fourteen",
	15: "Fifteen",
	16: "Sixteen",
	17: "Seventeen",
	18: "Eighteen",
	19: "Nineteen",
}
var tens = map[int]string{
	2: "Twenty",
	3: "Thirty",
	4: "Forty",
	5: "Fifty",
	6: "Sixty",
	7: "Seventy",
	8: "Eighty",
	9: "Ninety",
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	var r = doTranslate(num)
	return strings.Join(r, " ")
}

func doTranslate(n int) []string {
	if n == 0 {
		return []string{}
	} else if n <= 19 {
		return []string{to19[n]}
	} else if n <= 99 {
		var ss = []string{tens[n/10]}
		ss = append(ss, doTranslate(n%10)...)
		return ss
	} else if n <= 999 {
		var ss = make([]string, 0)
		var d = n / 100
		ss = append(ss, doTranslate(d)...)
		ss = append(ss, "Hundred")
		n -= d * 100
		ss = append(ss, doTranslate(n)...)
		return ss
	}
	var ss = make([]string, 0)
	for i := 30; i >= 3; i -= 3 {
		var (
			p = int(math.Pow(10.0, float64(i)))
			d = n / p
		)
		if d >= 1 {
			ss = append(ss, doTranslate(d)...)
			ss = append(ss, ds[i])
			n -= d * p
		}
	}
	ss = append(ss, doTranslate(n)...)
	return ss
}

func main() {
	var ts = []struct {
		i int
		e string
	}{
		{i: 0, e: "Zero"},
		{i: 1, e: "One"},
		{i: 9, e: "Nine"},
		{i: 10, e: "Ten"},
		{i: 19, e: "Nineteen"},
		{i: 20, e: "Twenty"},
		{i: 21, e: "Twenty One"},
		{i: 30, e: "Thirty"},
		{i: 31, e: "Thirty One"},
		{i: 90, e: "Ninety"},
		{i: 99, e: "Ninety Nine"},
		{i: 100, e: "One Hundred"},
		{i: 1000, e: "One Thousand"},
		{i: 2000, e: "Two Thousand"},
		{i: 2222, e: "Two Thousand Two Hundred Twenty Two"},
		{i: 100010001000, e: "One Hundred Billion Ten Million One Thousand"},
	}
	for _, t := range ts {
		var a = numberToWords(t.i)
		if t.e != a {
			fmt.Printf("numberToWords(%v) should have been %v but was %v\n", t.i, t.e, a)
		}
	}
}
