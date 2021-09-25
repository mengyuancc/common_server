package base_test

import (
	"fmt"
	"testing"
)


var (
	e = "a"
	s = "hello"
	z = true
)

func TestBase(t *testing.T) {
	defer fmt.Println("end1")
	defer fmt.Println("end2")
	var a int = 6
	var b int
	b = 8
	var c = 9
	d := 10
	fmt.Println(a, b, c, d, e, s, z)

	const PI float32 = 3.1415926
	fmt.Println(PI)

	type inter int
	var v inter = 3
	fmt.Println(v)
	var myFunc func(string)
	myFunc = func(name string) {
		fmt.Println(name)
	}
	myFunc("mengyuan")

	fmt.Println(sum(1,4,5,6,7,8))

	var p1 int = 2
	var p2 *int = &p1
	*p2 = 3
	fmt.Println(p1)

	w, p := 3, 4
	fmt.Println(switchs(&w, &p))
}

func sum(args ...int) int {
	ret := 0
	for _, v := range args {
		ret += v
	}
	return ret
}

func switchs(a, b *int) (int, int){
	return *b, *a
}



