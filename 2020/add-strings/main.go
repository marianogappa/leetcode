package main

import "fmt"

// Time: O(n)
// Space: O(n) or O(1) if inverted (solution space) doesn't count
func addStrings(num1 string, num2 string) string {
	inverted := []byte{}
	carry := 0
	for i := 0; i < len(num1) || i < len(num2); i++ {
		var n1, n2 int
		if i < len(num1) {
			n1 = int(num1[len(num1)-1-i] - '0')
		}
		if i < len(num2) {
			n2 = int(num2[len(num2)-1-i] - '0')
		}
		digit := (n1 + n2 + carry) % 10
		carry = (n1 + n2 + carry) / 10
		inverted = append(inverted, '0'+byte(digit))
	}
	if carry > 0 {
		inverted = append(inverted, '0'+byte(carry))
	}
	reverse(inverted)
	return string(inverted)
}

func reverse(bs []byte) {
	l := 0
	r := len(bs) - 1
	for l < r {
		bs[l], bs[r] = bs[r], bs[l]
		l++
		r--
	}
}

func main() {
	ts := []struct {
		num1     string
		num2     string
		expected string
	}{
		{
			num1:     "",
			num2:     "",
			expected: "",
		},
		{
			num1:     "123",
			num2:     "456",
			expected: "579",
		},
		{
			num1:     "999",
			num2:     "1",
			expected: "1000",
		},
	}
	for _, tc := range ts {
		actual := addStrings(tc.num1, tc.num2)
		if tc.expected != actual {
			fmt.Printf("For (%v,%v) expected %v but got %v\n", tc.num1, tc.num2, tc.expected, actual)
		}
	}
}
