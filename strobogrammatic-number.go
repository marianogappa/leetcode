package main

import "fmt"

func strobo(b1, b2 byte) bool {
	switch {
	case b1 == '8' && b2 == '8',
		b1 == '1' && b2 == '1',
		b1 == '0' && b2 == '0',
		b1 == '6' && b2 == '9',
		b1 == '9' && b2 == '6':
		return true
	}
	return false
}

func isStrobogrammatic(num string) bool {
	var i, j int
	j = len(num) - 1
	for i <= j {
		if !strobo(num[i], num[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println("69", isStrobogrammatic("69"))
	fmt.Println("88", isStrobogrammatic("88"))
	fmt.Println("818", isStrobogrammatic("818"))
	fmt.Println("8168", isStrobogrammatic("8168"))
	fmt.Println("81168", isStrobogrammatic("81168"))
	fmt.Println("89118", isStrobogrammatic("89118"))
	fmt.Println("1", isStrobogrammatic("1"))
	fmt.Println("4", isStrobogrammatic("4"))
}
