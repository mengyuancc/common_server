package common

import "fmt"

/**
 * 计算两个数的四则运算函数
 */
func Eval(a, b int, op string) (int, error) {
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

// 获取字符串最大不重复字串的长度
func lengthOfNonReportingSubstr(str string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i,ch := range []rune(str) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start +1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}
