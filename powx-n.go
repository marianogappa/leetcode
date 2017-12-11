package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	if n%2 == 0 {
		return myPow(x*x, n/2)
	}
	return x * myPow(x*x, n/2)
}

func main() {
	fmt.Println(myPow(0.00001, 2147483647))
	fmt.Println(myPow(2.00000, 10))
	fmt.Println(myPow(2.10000, 3))
}
