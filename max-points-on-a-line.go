package main

import "fmt"

type Point struct {
	X int
	Y int
}

type lf struct {
	m, b frac
}

type frac struct {
	n, d int
}

func (f1 frac) minus(f2 frac) frac {
	return frac{f1.n*f2.d - f2.n*f1.d, f1.d * f2.d}.simplify()
}

func (f1 frac) times(f2 frac) frac {
	return frac{f1.n * f2.n, f1.d * f2.d}.simplify()
}

func (f1 frac) div(f2 frac) frac {
	return frac{f1.n * f2.d, f1.d * f2.n}.simplify()
}

func (f frac) simplify() frac {
	if f.n < 0 && f.d < 0 {
		f.n, f.d = -f.n, -f.d
	}
	if f.n == f.d {
		return frac{1, 1}
	}
	if f.n == -f.d {
		return frac{-1, 1}
	}
	if f.n == 0 {
		return frac{0, 1}
	}
	if abs(f.n-f.d) == 1 {
		return f
	}
	for i := 2; i <= min(abs(f.n), abs(f.d)); i++ {
		for f.n%i == 0 && f.d%i == 0 {
			f.n /= i
			f.d /= i
		}
	}
	return f
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func maxPoints(points []Point) int {
	var (
		lfs = make(map[lf]map[int]struct{}, 0)
		xs  = make(map[int]int, 0)
		max = 0
	)
	for i, p1 := range points {
		xs[p1.X]++
		var x = xs[p1.X]
		if x > max {
			max = x
		}
		for j, p2 := range points {
			if i < j && p1.X != p2.X {
				var (
					m = frac{p1.Y - p2.Y, 1}.div(frac{p1.X - p2.X, 1})
					b = frac{p1.Y, 1}.minus(frac{p1.X, 1}.times(m))
					l = lf{m, b}
				)
				var mp = make(map[int]struct{}, 0)
				if _, ok := lfs[l]; ok {
					mp = lfs[l]
				}
				mp[i] = struct{}{}
				mp[j] = struct{}{}
				lfs[l] = mp
				if len(mp) > max {
					max = len(mp)
				}
			}
		}
	}
	return max
}

func main() {
	var ts = []struct {
		points []Point
		e      int
	}{
		{
			points: []Point{
				{0, 2}, {1, 2}, {2, 2},
			},
			e: 3,
		},

		{
			points: []Point{
				{1, 0}, {1, 1}, {1, 2},
			},
			e: 3,
		},
		{
			points: []Point{
				{0, 1}, {1, 2}, {2, 0},
			},
			e: 2,
		},
		{
			points: []Point{
				{0, 2}, {1, 1}, {2, 0},
			},
			e: 3,
		},
		{
			points: []Point{
				{1, 1}, {1, 1}, {2, 3},
			},
			e: 3,
		},
		{
			points: []Point{
				{560, 248}, {0, 16}, {30, 250}, {950, 187}, {630, 277}, {950, 187}, {-212, -268}, {-287, -222}, {53, 37}, {-280, -100}, {-1, -14}, {-5, 4}, {-35, -387}, {-95, 11}, {-70, -13}, {-700, -274}, {-95, 11}, {-2, -33}, {3, 62}, {-4, -47}, {106, 98}, {-7, -65}, {-8, -71}, {-8, -147}, {5, 5}, {-5, -90}, {-420, -158}, {-420, -158}, {-350, -129}, {-475, -53}, {-4, -47}, {-380, -37}, {0, -24}, {35, 299}, {-8, -71}, {-2, -6}, {8, 25}, {6, 13}, {-106, -146}, {53, 37}, {-7, -128}, {-5, -1}, {-318, -390}, {-15, -191}, {-665, -85}, {318, 342}, {7, 138}, {-570, -69}, {-9, -4}, {0, -9}, {1, -7}, {-51, 23}, {4, 1}, {-7, 5}, {-280, -100}, {700, 306}, {0, -23}, {-7, -4}, {-246, -184}, {350, 161}, {-424, -512}, {35, 299}, {0, -24}, {-140, -42}, {-760, -101}, {-9, -9}, {140, 74}, {-285, -21}, {-350, -129}, {-6, 9}, {-630, -245}, {700, 306}, {1, -17}, {0, 16}, {-70, -13}, {1, 24}, {-328, -260}, {-34, 26}, {7, -5}, {-371, -451}, {-570, -69}, {0, 27}, {-7, -65}, {-9, -166}, {-475, -53}, {-68, 20}, {210, 103}, {700, 306}, {7, -6}, {-3, -52}, {-106, -146}, {560, 248}, {10, 6}, {6, 119}, {0, 2}, {-41, 6}, {7, 19}, {30, 250},
			},
			e: 22,
		},
		{
			points: []Point{
				{0, 0}, {-1, -1}, {2, 2},
			},
			e: 3,
		},
		{
			points: []Point{
				{2, 3}, {3, 3}, {-5, 3},
			},
			e: 3,
		},
		{
			points: []Point{
				{3, 10}, {0, 2}, {0, 2}, {3, 10},
			},
			e: 4,
		},
		{
			points: []Point{
				{0, 0}, {1921151, 1921150}, {1921152, 1921151},
			},
			e: 2,
		},
	}
	for _, t := range ts {
		var a = maxPoints(t.points)
		if a != t.e {
			fmt.Printf("maxPoints(%v) should have been %v but was %v\n", t.points, t.e, a)
		}
	}
}
