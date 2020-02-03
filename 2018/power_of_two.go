package main

import "fmt"

func isPowerOfTwo(n int) bool {
	return n != 0 && n&(n-1) == 0
}

func main() {
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(17))
}
