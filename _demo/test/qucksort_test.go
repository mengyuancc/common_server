package quick_test

import "testing"

func QuickSort(a []int) []int{
	if len(a) <=1 {
		return a
	}
	var left []int
	var right []int
	flag := a[0]
	for i:=1; i<len(a); i++ {
		if a[i] <= flag {
			left = append(left, a[i])
		} else {
			right = append(right, a[i])
		}
	}
	left = QuickSort(left)
	right = QuickSort(right)

	b := []int{flag}
	return append(append(left, b...), right...)
}

func TestQuick(t *testing.T) {
	a := []int{3,5,2,1,4,100,6}
	b := QuickSort(a)
	t.Log(b)
}