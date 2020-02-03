package main

import "fmt"

func lengthLongestPath(input string) int {
	var (
		cur, curDirLen, max, tokLen, tabs int
		isFile                            bool
		curDir                            = make([]int, 0)
	)
	for {
		tokLen, isFile, cur, tabs = readTok(input, cur)
		if isFile {
			fileLen := curDirLen + len(curDir) + tokLen
			if fileLen > max {
				max = fileLen
			}
		} else { // isDirectory
			if tabs > len(curDir) {
				curDir = append(curDir, tokLen)
				curDirLen += tokLen
			}
		}

		if tabs <= len(curDir)-1 {
			remove := len(curDir) - tabs
			for i := 0; i < remove; i++ {
				curDirLen -= curDir[len(curDir)-1-i]
			}
			curDir = curDir[:len(curDir)-remove]
		}

		if cur >= len(input)-1 {
			break
		}
	}
	return max
}

func readTok(input string, cur int) (tokLen int, isFile bool, newCur int, tabs int) {
	from := cur
	for cur < len(input) && input[cur] != '\n' {
		if input[cur] == '.' {
			isFile = true
		}
		cur++
	}
	cur++

	for cur < len(input) {
		if input[cur] == '\t' {
			tabs++
			cur++
		} else {
			break
		}
	}
	return cur - from - tabs - 1, isFile, cur, tabs
}

func main() {
	fss := []string{
		"",
		"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext",
		"dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext",
		"dir\n    file.txt",
		"a\n\tb1\n\t\tf1.txt\n\taaaaa\n\t\tf2.txt",
		"a\n\tb.txt\na2\n\tb2.txt",
	}
	for _, fs := range fss {
		fmt.Printf("--------------\nfor \n%v\nit's %v\n--------------\n", fs, lengthLongestPath(fs))
	}
}
