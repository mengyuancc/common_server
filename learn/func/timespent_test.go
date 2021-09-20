package func_test

import (
	"fmt"
	"testing"
	"time"
)

type option func(op int) int

func timeSpent(inner option) option{
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:",time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int{
	time.Sleep(time.Second*1)
	return op
}

func Sum(ops ...int) int {
	var ret int = 0
	for _,v := range ops {
		ret = ret + v
	}
	return ret
}

func TestFun(t *testing.T){
	//tsSf := timeSpent(slowFun)
	//tsSf(10)

	t.Log(Sum(1,3,5,6,7,8,9,4))
}

func mapTest(m map[string]string) {
	if _,ok := m["name"];ok {
		m["name"] = "xiaoming"
	} else {
		fmt.Println("error")
	}
}

func TestMapChange(t *testing.T) {
	var m = map[string]string{
		"name":"mengyuan",
	}
	mapTest(m)
	t.Log(m)
}
