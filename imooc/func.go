package main

import "fmt"

/**
 * 计算两个数的四则运算函数
 */
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation:%s", op)
	}
}
func main() {
	if c, err := eval(1, 3, "V"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c)
	}
}
