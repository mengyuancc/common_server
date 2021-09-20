package common

import (
	"fmt"
	"testing"
)

/**
 * 普通的测试
 */
func TestEval(t *testing.T) {
	r, err := Eval(1,2, "+")
	if err != nil{
		t.Error("error")
	}
	fmt.Println(r)
}

// benchmark 测试
func BenchmarkEval(b *testing.B) {
	s := "这是一个测试的字符串字符串"
	ans := 10
	for i:=0; i<b.N; i++ {
		actual := lengthOfNonReportingSubstr(s)
		if actual != ans {
			b.Errorf("got %d for input %s: expected %d", actual, s, ans)
		}
	}
}
