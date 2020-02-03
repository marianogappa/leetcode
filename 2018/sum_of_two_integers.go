package main

import "fmt"

func getSum(a int, b int) int {
	for a&b != 0 {
		a, b = a^b, a&b<<1
	}
	return a | b
}

func main() {
	fmt.Println(getSum(4, 50))
	fmt.Println(getSum(4, 5))
}
