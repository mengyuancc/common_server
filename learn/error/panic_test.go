package error_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicDemo(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			//记录log或者恢复 let it crash
			t.Log("show err",err)
		}
	}()

	fmt.Println("start")

	panic(errors.New("错误，异常退出"))

	fmt.Println("end")
}