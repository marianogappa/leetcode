package main

import "fmt"

func shortestDistance(words []string, word1 string, word2 string) int {
	var (
		minD         = 1<<31 - 1
		start        = -1
		startIsWord1 bool
	)
	for i, w := range words {
		switch w {
		case word1:
			if start != -1 && !startIsWord1 {
				if i-start < minD {
					minD = i - start
				}
			}
			startIsWord1 = true
			start = i
		case word2:
			if start != -1 && startIsWord1 {
				if i-start < minD {
					minD = i - start
				}
			}
			startIsWord1 = false
			start = i
		}
	}
	return minD
}

func main() {
	fmt.Println(shortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "coding", "practice"))
	fmt.Println(shortestDistance([]string{"practice", "makes", "perfect", "coding", "makes"}, "makes", "coding"))
}
