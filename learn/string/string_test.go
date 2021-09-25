package string_test

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

// 计算字符串中最长的不重复的字符串的长度
func lengthOfRepeatString(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func TestStringInit(t *testing.T) {
	var s string = "问君能有几多愁，恰似已经春水向东流"
	// 获取字符串的字数
	num := len([]rune(s))
	t.Log(num)
	t.Log("rune count:", utf8.RuneCountInString(s))
	t.Log(num)
	// 获取字符串的字节数
	num = len(s)
	// 遍历字符串
	// 输出utf-8码
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	// 输出unicode码
	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()
	// 输出char
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	// 字符串操作
	// Fields,Split,Join
	// Contains,Index
	// ToLower,ToUpper
	// Trim,TrimRight,TrimLeft
	fmt.Println()
	a := strings.Split(s, "")
	fmt.Println(a)
	for i, ch := range a {
		fmt.Print("(",i, ch, ")")
	}
}