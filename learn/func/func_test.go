package func_test

import (
	"testing"
)

// 函数可以赋值给变量
// 这是函数式编程的开始，也是比较重要的一点
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

// 闭包实现累加
func adder() func(int) int {
	// 自由变量
	sum := 0
	// 函数体
	return func(v int) int { // v局部变量
		sum += v
		return sum
	}
}
func TestAdd(t *testing.T) {
	a := adder()
	for i := 0; i <  10; i++ {
		t.Logf("0 + 1 + ... + %d = %d\n", i, a(i))
	}
}

