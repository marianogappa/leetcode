package main

import "fmt"

func compareVersion(version1 string, version2 string) int {
	var (
		v1 = readVersion(version1)
		v2 = readVersion(version2)
		i  int
	)
	for i < len(v1) || i < len(v2) {
		if i >= len(v1) && i < len(v2) {
			return -1
		} else if i < len(v1) && i >= len(v2) {
			return 1
		} else if v1[i] < v2[i] {
			return -1
		} else if v1[i] > v2[i] {
			return 1
		}
		i++
	}
	return 0
}

func readVersion(s string) []int {
	var (
		rs = make([]int, 0)
		ns = make([]byte, 0)
	)

	for i := 0; i < len(s); i++ {
		if s[i] == '.' && len(ns) > 0 {
			rs = append(rs, toInt(ns))
			ns = make([]byte, 0)
		} else {
			ns = append(ns, s[i])
		}
	}
	if len(ns) > 0 {
		rs = append(rs, toInt(ns))
	}

	// Remove zeroes from the right
	var j int
	for j = len(rs) - 1; j >= 0; j-- {
		if rs[j] != 0 {
			break
		}
	}

	return rs[:j+1]
}

func toInt(ns []byte) int {
	var (
		n = 0
		d = 1
	)
	for i := len(ns) - 1; i >= 0; i-- {
		n += int(ns[i]-'0') * d
		d *= 10
	}
	return n
}

func main() {
	fmt.Println(compareVersion("1", "1.0.0.0.0.0") == 0)
	fmt.Println(compareVersion("1.0", "1.0.0") == 0)
	fmt.Println(compareVersion("", "") == 0)
	fmt.Println(compareVersion("1", "2") == -1)
	fmt.Println(compareVersion("2", "1") == 1)
	fmt.Println(compareVersion("1.2", "1.2") == 0)
	fmt.Println(compareVersion("1.1", "1.2") == -1)
	fmt.Println(compareVersion("1.1.1", "1.1.2") == -1)
	fmt.Println(compareVersion("1.1.2", "1.1.1") == 1)
	fmt.Println(compareVersion("1.1111.2", "1.1111.2") == 0)
	fmt.Println(compareVersion("1.1112.2", "1.1111.2") == 1)
	fmt.Println(compareVersion("1.1112.2", "1.1114.2") == -1)
}
