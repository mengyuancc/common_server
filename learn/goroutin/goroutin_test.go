package goroutin_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

var wg sync.WaitGroup
func worker(i int) {
	defer wg.Done()
	fmt.Printf("worker %d start doing \n", i)
}

func TestWorker(t *testing.T) {
	for i:=0; i<10000; i++ {
		wg.Add(1)
		go worker(i)
	}
	wg.Wait()
}

func TestRunTime(t *testing.T) {
	go func(s string) {
		for i:=0; i<100; i++ {
			fmt.Println(s)
		}
	}("world")
	for i:=0; i<100; i++ {
		// 切换一下
		runtime.Gosched()
		fmt.Println("hello")
	}
}
