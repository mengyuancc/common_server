package error_test

import (
	"errors"
	"testing"
)
var LessThanTwoError = errors.New("n should be not less than 2")
var MoreThanTenError = errors.New("n should be not more than 100")
func GetFibonacci(n int) ([]int,error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, MoreThanTenError
	}
	fib := []int{1,1}
	for i := 2; i<n;i++ {
		fib = append(fib, fib[i-2]+fib[i-1])
	}
	return fib, nil
}

func TestGetFibonacci(t *testing.T)  {
	if data,err := GetFibonacci(101); err != nil {
		if err == LessThanTwoError {
			t.Log("It is less")
		}
		if err == MoreThanTenError {
			t.Log("It is More")
		}
	} else {
		t.Log(data)
	}
}