package goroutine_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// goroutine
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
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)
}

// channel
func recv(c chan int) {
	c <- 3
	fmt.Println("发送成功")
}

func TestChannelDemo(t *testing.T) {
	ch := make(chan int)
	go recv(ch)
	<-ch
	t.Log("接受成功")
}

// 优雅的从channel中循环取值
func TestChannelDemo2(t *testing.T) {
	c := make(chan int)
	go func() {
		for i:=0; i< 10; i++ {
			c <- i
			fmt.Println("发送", i)
		}
		close(c)
	}()
	/*for {
		if data,ok := <-c; ok {
			t.Log(data)
		} else {
			break
		}
	}*/
	for v := range c {
		fmt.Println("接收", v)
		// t.Log(v)
	}
}

