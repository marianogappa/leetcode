package main

import "fmt"

// Time: O(n) for Count & O(1) for Add
// Space: O(n) save all points in a hashmap
//
// Pretty straightforward. Only trick is how to find squares.
// Two points can form a square only if |x1-x2| = |y1-y2|. The other points
// share one coordinate from one point, and one from the other.
//
// The only other mini-trick is that there might be many points on the
// same coordinates, so keep count of how many points are in each (x,y),
// and multiply the counts to answer how many squares can be formed.
//
// Last mini-trick: area must be positive, so don't try to form squares
// using the same two points as starting point.
type DetectSquares struct {
	points map[int]map[int]int
}

func Constructor() DetectSquares {
	return DetectSquares{points: map[int]map[int]int{}}
}

func (this *DetectSquares) Add(point []int) {
	if _, ok := this.points[point[0]]; !ok {
		this.points[point[0]] = map[int]int{}
	}
	this.points[point[0]][point[1]]++
}

func (this *DetectSquares) countPointsAt(x, y int) int {
	_, ok := this.points[y]
	if !ok {
		return 0
	}
	return this.points[y][x]
}

func (this *DetectSquares) Count(point []int) int {
	count := 0
	for y := range this.points {
		for x, c1 := range this.points[y] {
			if y == point[0] && x == point[1] {
				continue
			}
			if abs(y-point[0]) != abs(x-point[1]) {
				continue
			}
			c2 := this.countPointsAt(point[1], y)
			c3 := this.countPointsAt(x, point[0])
			count += c1 * c2 * c3
		}
	}
	return count
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */

func main() {
	obj := Constructor()
	pobj := &obj
	pobj.Add([]int{3, 10})
	pobj.Add([]int{11, 2})
	pobj.Add([]int{3, 2})
	fmt.Println(pobj.Count([]int{11, 10}) == 1)
	fmt.Println(pobj.Count([]int{14, 8}) == 0)
	pobj.Add([]int{11, 2})
	fmt.Println(pobj.Count([]int{11, 10}) == 2)
}
