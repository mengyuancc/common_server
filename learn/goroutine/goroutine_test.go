package goroutine_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func singleWork(t *testing.T) {
	t.Log("work hello Goroutine")
}

// 启动单个goroutine
func TestSingleGoRoutine(t *testing.T) {
	go singleWork(t)
	t.Log("main goroutine done")
	// time.Sleep(time.Second) 强制主协程等待子协程
}

var wg sync.WaitGroup
func worker(i int) {
	defer wg.Done()
	fmt.Printf("worker %d start doing \n", i)
}

// 启动多个goroutine （使用sync.WaitGroup来实现goroutine的同步）
func TestWorker(t *testing.T) {
	for i:=0; i<10000; i++ {
		wg.Add(1)
		go worker(i)
	}
	wg.Wait()
}

// runtime
func TestRunTime(t *testing.T) {
	go func(s string) {
		for i:=0; i<10; i++ {
			fmt.Println(s)
		}
	}("world")
	for i:=0; i<10; i++ {
		// 切换一下
		if i == 5 {
			runtime.Gosched()
		}
		fmt.Println("hello")
	}
}

func a() {
	for i := 1; i < 4; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 4; i++ {
		fmt.Println("B:", i)
	}
}

func TestRuntimeGOMax(t *testing.T) {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}

