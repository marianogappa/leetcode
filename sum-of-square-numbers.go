package main

import "fmt"

func judgeSquareSum(c int) bool {
	if c <= 1 {
		return true
	}
	var a = 0
	for a*a <= c {
		if sqrtExists(c - a*a) {
			return true
		}
		a++
	}
	return false
}

func sqrtExists(n int) bool {
	var r = n
	for r*r > n {
		r = (r + n/r) / 2
	}
	return r*r == n
}

//integer √c-a*a exists
func main() {
	for i := 0; i < 1000; i++ {
		fmt.Print(judgeSquareSum(i))
	}
}

/*
This solution is much better; lower and higher possible side of the triangle

func judgeSquareSum(c int) bool {
    if c < 0 {
    return false
  }
  root := int(math.Sqrt(float64(c))) + 1
  low, high := 0, root
  for low <= high {
    cur := low*low + high*high
    if cur < c {
      low++
    } else if cur > c {
      high--
    } else {
      return true
    }
  }
  return false
}
*/
