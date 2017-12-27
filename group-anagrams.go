package main

import "fmt"

func qsort(a []byte) {
	if len(a) < 2 {
		return
	}
	left, right := 0, len(a)-1
	pivotIndex := len(a) / 2
	a[pivotIndex], a[right] = a[right], a[pivotIndex]
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]
	qsort(a[:left])
	qsort(a[left+1:])
	return
}

func isAnagram(s1, s2 string) bool {
	var b byte
	for i := 0; i < len(s1); i++ {
		b ^= s1[i] ^ s2[i]
	}
	if b != 0 {
		return false
	}
	var (
		b1 = []byte(s1)
		b2 = []byte(s2)
	)
	qsort(b1)
	qsort(b2)
	return string(b1) == string(b2)
}

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	var mAns = make(map[int][][]string, 0)
	for i := 0; i < len(strs); i++ {
		var found bool
		for j, a := range mAns[len(strs[i])] {
			if isAnagram(strs[i], a[0]) {
				var ss = mAns[len(strs[i])]
				ss[j] = append(ss[j], strs[i])
				mAns[len(strs[i])] = ss
				found = true
				break
			}
		}
		if !found {
			var s = mAns[len(strs[i])]
			s = append(s, []string{strs[i]})
			mAns[len(strs[i])] = s
		}
	}
	var ss = make([][]string, 0)
	for _, a := range mAns {
		for _, b := range a {
			ss = append(ss, b)
		}
	}
	return ss
}

func main() {
	var ts = []struct {
		i []string
	}{
		{i: []string{""}},
		{i: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
		{i: []string{"tao", "pit", "cam", "aid", "pro", "dog"}},
		{i: []string{"tao", "ato", "cam", "aid", "pro", "dog"}},
		{i: []string{"tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao", "tao"}},
	}
	for _, t := range ts {
		fmt.Printf("for %v got %v\n", t.i, groupAnagrams(t.i))
	}
}
