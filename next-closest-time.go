package main

import "fmt"

func nextClosestTime(time string) string {
	var nums = make(map[byte]struct{})
	nums[time[0]] = struct{}{}
	nums[time[1]] = struct{}{}
	nums[time[3]] = struct{}{}
	nums[time[4]] = struct{}{}
	var min, _ = traverse(nums, atom(string([]byte{time[0], time[1], time[3], time[4]})), 1<<31-1, "", "")
	if min == "" {
		return time
	}
	return min[:2] + ":" + min[2:]
}

func validSoFar(s string) bool {
	switch {
	case len(s) == 1:
		return s[0] >= '0' && s[0] <= '2'
	case len(s) == 2 && s[0] == '2':
		return s[1] <= '4'
	case len(s) == 3:
		return s[2] <= '5'
	}
	return true
}

func atom(s string) int {
	b1, b2, b3, b4 := int(s[0]-'0'), int(s[1]-'0'), int(s[2]-'0'), int(s[3]-'0')
	return (b1*10+b2)*60 + b3*10 + b4
}

func dist(f, t int) int {
	if t >= f {
		return t - f
	}
	return 24*60 - f + t
}

func traverse(nums map[byte]struct{}, floor, minDist int, min string, cur string) (string, int) {
	if len(cur) < 4 && !validSoFar(cur) {
		return min, minDist
	}
	if len(cur) == 4 {
		var d = dist(floor, atom(cur))
		if d < minDist && d > 0 {
			return cur, d
		}
		return min, minDist
	}
	for n := range nums {
		min, minDist = traverse(nums, floor, minDist, min, cur+string(n))
	}
	return min, minDist
}

func main() {
	var ts = []struct {
		i string
		e string
	}{
		{i: "19:34", e: "19:39"},
		{i: "23:59", e: "22:22"},
		{i: "00:00", e: "00:00"},
		{i: "11:11", e: "11:11"},
	}
	for _, t := range ts {
		var a = nextClosestTime(t.i)
		if t.e != a {
			fmt.Printf("nextClosestTime(%v) should have been %v but was %v\n", t.i, t.e, a)
		}
	}
}
