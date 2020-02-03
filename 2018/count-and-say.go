package main

import "fmt"

func countAndSay(n int) string {
	if n <= 0 {
		return ""
	}
	var s = "1"
	for i := 2; i <= n; i++ {
		var (
			tmpS  string
			curB  byte
			count int
		)
		for j := 0; j < len(s); j++ {
			if count == 0 {
				curB = s[j]
				count = 1
			} else {
				if s[j] == curB {
					count++
				} else {
					tmpS += fmt.Sprintf("%v%v", count, string(curB))
					curB = s[j]
					count = 1
				}
			}
		}
		if count > 0 {
			tmpS += fmt.Sprintf("%v%v", count, string(curB))
		}
		s = tmpS
	}
	return s
}

func main() {
	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))
	fmt.Println(countAndSay(6))
	fmt.Println(countAndSay(7))
	fmt.Println(countAndSay(8))
	fmt.Println(countAndSay(9))
	fmt.Println(countAndSay(10))
	fmt.Println(countAndSay(11))
	fmt.Println(countAndSay(12))
	fmt.Println(countAndSay(13))
	fmt.Println(countAndSay(14))
	fmt.Println(countAndSay(15))
	fmt.Println(countAndSay(16))
	fmt.Println(countAndSay(17))
	fmt.Println(countAndSay(18))
	fmt.Println(countAndSay(19))
	fmt.Println(countAndSay(20))
	fmt.Println(countAndSay(21))
	fmt.Println(countAndSay(22))
	fmt.Println(countAndSay(23))
	fmt.Println(countAndSay(24))
	fmt.Println(countAndSay(25))
	fmt.Println(countAndSay(26))
	fmt.Println(countAndSay(27))
	fmt.Println(countAndSay(28))
	fmt.Println(countAndSay(29))
	fmt.Println(countAndSay(30))
	fmt.Println(countAndSay(31))
	fmt.Println(countAndSay(32))
	fmt.Println(countAndSay(33))
	fmt.Println(countAndSay(34))
	fmt.Println(countAndSay(35))
}
