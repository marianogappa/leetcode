package main

import "fmt"

func reverseWords(str []byte) {
	var (
		i = 0
		j = len(str) - 1
	)
	for i < j {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}

	var start = 0
	for i := 0; i < len(str); i++ {
		if (str[i] == ' ' || i == len(str)-1) && i-start > 0 {
			var (
				k = start
				l = i - 1
			)
			if i == len(str)-1 {
				l = i
			}
			for k < l {
				str[k], str[l] = str[l], str[k]
				k++
				l--
			}
			start = i + 1
		}
	}
}

func main() {
	var a = []byte("the sky is blue")
	reverseWords(a)
	fmt.Println(string(a))
}
