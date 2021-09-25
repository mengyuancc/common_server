package error

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"testing"
)

func tryDefer() {
	// defer fmt.Println(1)
	// fmt.Println(2)
	for i:=0; i<30; i++ {
		defer fmt.Println(i)
		if i == 10 {
			panic("end")
		}
	}
}

func Fibonacci() func() int {
	var a,b = 0,1
	return func() int {
		a, b = b, a + b
		return a
	}
}

func writeFile(filename string) {
	/*file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}*/
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil  {
		// 自定义error
		err = errors.New("this is custom error")
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf(" %s, %s, %s\n",
				pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()
	// 内存buffer
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := Fibonacci()
	for i:=0; i<30; i++ {
		// 写入buffer
		fmt.Fprintln(writer, f())
	}
}
func TestDefer(t *testing.T) {
	// tryDefer()
	writeFile("fib.txt")
}
