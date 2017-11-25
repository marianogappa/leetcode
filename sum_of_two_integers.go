package main

import "fmt"

func getSum(a int, b int) int {
	nonCarries, carries := a^b, a&b<<1

	for nonCarries&carries != 0 {
		nonCarries, carries = nonCarries^carries, nonCarries&carries<<1
	}
	return nonCarries | carries
}

func main() {
	fmt.Println(getSum(4, 50))
	fmt.Println(getSum(4, 5))
}
