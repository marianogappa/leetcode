package main

import (
	"fmt"
	"strings"
)

func numberToWords(num int) string {
	if num <= 999 {
		return upTo999ToWords(num)
	}

	triplets := []int{}
	for num != 0 {
		triplets = append(triplets, num%1000)
		num /= 1000
	}

	english := []string{}
	for i := len(triplets) - 1; i >= 0; i-- {
		if triplets[i] == 0 {
			continue
		}
		english = append(english, upTo999ToWords(triplets[i]))
		if i > 0 {
			english = append(english, tripletsEnglish[i])
		}
	}
	return strings.Join(english, " ")
}

var (
	firstDigitEnglish = map[int]string{
		0: "Zero",
		1: "One",
		2: "Two",
		3: "Three",
		4: "Four",
		5: "Five",
		6: "Six",
		7: "Seven",
		8: "Eight",
		9: "Nine",
	}
	secondDigitEnglish = map[int]string{
		2: "Twenty",
		3: "Thirty",
		4: "Forty",
		5: "Fifty",
		6: "Sixty",
		7: "Seventy",
		8: "Eighty",
		9: "Ninety",
	}
	tripletsEnglish = map[int]string{
		1: "Thousand",
		2: "Million",
		3: "Billion",
		4: "Trillion",
		5: "Quadrillion",
		6: "Quintillion",
		7: "Sixtillion",
		8: "Septillion",
		9: "Octillion",
	}
)

func upTo999ToWords(num int) string {
	if num < 100 {
		return upTo99ToWords(num)
	}
	rest := num % 100
	howManyHundreds := num / 100

	if rest == 0 {
		return fmt.Sprintf("%v Hundred", firstDigitEnglish[howManyHundreds])
	}
	return fmt.Sprintf("%v Hundred %v", firstDigitEnglish[howManyHundreds], upTo99ToWords(rest))
}

func upTo99ToWords(num int) string {
	if num < 10 {
		return firstDigitEnglish[num]
	}

	switch num {
	case 10:
		return "Ten"
	case 11:
		return "Eleven"
	case 12:
		return "Twelve"
	case 13:
		return "Thirteen"
	case 14:
		return "Fourteen"
	case 15:
		return "Fifteen"
	case 16:
		return "Sixteen"
	case 17:
		return "Seventeen"
	case 18:
		return "Eighteen"
	case 19:
		return "Nineteen"
	}

	firstDigit := num % 10
	secondDigit := (num / 10) % 10

	if firstDigit == 0 {
		return secondDigitEnglish[secondDigit]
	}
	return fmt.Sprintf("%v %v", secondDigitEnglish[secondDigit], firstDigitEnglish[firstDigit])
}

func main() {
	ts := []struct {
		input    int
		expected string
	}{
		{
			input:    0,
			expected: "Zero",
		},
		{
			input:    123,
			expected: "One Hundred Twenty Three",
		},
		{
			input:    12345,
			expected: "Twelve Thousand Three Hundred Forty Five",
		},
		{
			input:    1234567,
			expected: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven",
		},
		{
			input:    1234567891,
			expected: "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One",
		},
		{
			input:    1000,
			expected: "One Thousand",
		},
		{
			input:    3000000,
			expected: "Three Million",
		},
	}
	for _, tc := range ts {
		actual := numberToWords(tc.input)
		if tc.expected != actual {
			fmt.Printf("For %v expected %v but got %v\n", tc.input, tc.expected, actual)
		}
	}
}
