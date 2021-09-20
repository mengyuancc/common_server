package func_test

import "testing"

func abs(num int) int {
	if num > 0 {
		return num
	} else {
		return -num
	}
}

var new = func(num int) int{
	if num > 0 {
		return num
	} else {
		return -num
	}
}

func TestCalculate(t *testing.T) {
	a := abs(-6)
	t.Log(a)
	t.Log(new(-3))
}
