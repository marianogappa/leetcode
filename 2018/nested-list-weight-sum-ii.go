package main

type NestedInteger struct {
}

func (n NestedInteger) GetInteger() int           { return 0 }
func (n NestedInteger) IsInteger() bool           { return true }
func (n NestedInteger) GetList() []*NestedInteger { return nil }

func depthSumInverse(nestedList []*NestedInteger) int {
	var (
		ss = make([]int, 0)
		s  = 0
	)
	iterate(nestedList, &ss, 0)
	for i, ps := range ss {
		s += ps * (len(ss) - i)
	}
	return s
}

func iterate(ns []*NestedInteger, ss *[]int, index int) {
	if len(*ss) == index {
		*ss = append(*ss, 0)
	}
	for _, n := range ns {
		if n.IsInteger() {
			(*ss)[index] += n.GetInteger()
		} else {
			iterate(n.GetList(), ss, index+1)
		}
	}
}

func main() {

}
