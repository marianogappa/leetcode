package main

type WordDistance struct {
	ws map[string][]int
}

func Constructor(words []string) WordDistance {
	var ws = make(map[string][]int, 0)
	for i, w := range words {
		ws[w] = append(ws[w], i)
	}
	return WordDistance{ws}
}

func (this *WordDistance) Shortest(word1 string, word2 string) int {
	var (
		d1     = this.ws[word1]
		d2     = this.ws[word2]
		minD   = 1<<31 - 1
		i1, i2 int
	)
	for minD > 1 && i1 < len(d1) && i2 < len(d2) {
		minD = min(minD, abs(d1[i1]-d2[i2]))
		if d1[i1] < d2[i2] {
			i1++
		} else {
			i2++
		}
	}
	return minD
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Shortest(word1,word2);
 */

func main() {

}
