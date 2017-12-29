package main

import "fmt"

type MovingAverage struct {
	l, i int
	vs   []int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{l: size, i: 0, vs: make([]int, 0, size)}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.i > len(this.vs)-1 {
		this.vs = append(this.vs, val)
	} else {
		this.vs[this.i] = val
	}
	this.i++
	this.i = this.i % this.l
	var sum int
	for _, v := range this.vs {
		sum += v
	}
	return float64(sum) / float64(len(this.vs))
}

func main() {
	var (
		ma = Constructor(3)
		p1 = ma.Next(1)
		p2 = ma.Next(10)
		p3 = ma.Next(3)
		p4 = ma.Next(5)
	)
	fmt.Println(p1, p2, p3, p4)
}
