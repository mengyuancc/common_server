package goroutine_test

import (
	"fmt"
	"testing"
)

// channel
func TestChannelInit(t *testing.T) {
	// 定义channel
	var ch chan int
	// 创建channel 初始化channel只有在初始化完成后才能使用
	// ch = make(chan int)
	ch = make(chan int, 1) //也可以设置缓冲大小
	t.Log(ch)
}

// channel操作
func TestChannelUse(t *testing.T){
	//var ch chan int
	//ch := make(chan int)
	//ch <- 2  // 发送
	//<-ch  // 接收
	//close(ch) //关闭
}

func recv(c chan int, t *testing.T) {
	c <- 3
	t.Log("发送成功")
}

// 无缓冲的通道 在没有设置缓冲区的时候需要 发送者 也需要接收者 需要开启两个协程
func TestNoCacheDemo(t *testing.T) {
	ch := make(chan int)
	// go recv(ch, t)
	close(ch)
	a := <-ch
	t.Log(a, "接受成功")
}

// 有缓冲区的通道
func TestCacheDemo(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 1
	t.Log(ch)
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

