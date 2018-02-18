package main

import "fmt"

func anagramMappings(A []int, B []int) []int {
	var (
		m = make(map[int][]int, len(B))
		r = make([]int, len(A))
	)
	for i, n := range B {
		m[n] = append(m[n], i)
	}
	for i, v := range A {
		r[i] = m[v][len(m[v])-1]
		m[v] = m[v][:len(m[v])-1]
	}
	return r
}

func main() {
	fmt.Println(anagramMappings([]int{12, 28, 46, 32, 50}, []int{50, 12, 32, 46, 28}))
	fmt.Println(anagramMappings([]int{12, 12, 12}, []int{12, 12, 12}))
	fmt.Println(anagramMappings([]int{12}, []int{12}))
}
