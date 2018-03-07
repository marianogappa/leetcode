package main

import "fmt"

func isPerfectSquare(num int) bool {
	var r = num
	for r*r > num {
		r = (r + num/r) / 2
	}
	return r*r == num
}

func main() {
	for i := 0; i <= 1000000; i++ {
		if isPerfectSquare(i) {
			fmt.Println(i)
		}
	}
}
