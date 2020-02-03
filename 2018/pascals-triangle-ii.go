package main

import "fmt"

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	var ns = []int{1, 2, 1}
	for i := 3; i <= rowIndex; i++ {
		var nns = []int{1}
		for j := 1; j < len(ns); j++ {
			nns = append(nns, ns[j-1]+ns[j])
		}
		nns = append(nns, 1)
		nns, ns = ns, nns
	}
	return ns
}

func main() {
	fmt.Println(getRow(0))
	fmt.Println(getRow(1))
	fmt.Println(getRow(2))
	fmt.Println(getRow(3))
	fmt.Println(getRow(4))
	fmt.Println(getRow(5))
	fmt.Println(getRow(100))
}
