package goroutine

import (
	"testing"
	"time"
)

// 空的channel 关闭成功后会返回零值
func TestChannelClose(t *testing.T) {
	ch := make(chan int)
	t.Log(ch)
	go func() {
		time.Sleep(3*time.Second)
		close(ch)
	}()
	select {
	case <-ch:
		t.Log("channel close")
	}
}

func TestSelectDemo(t *testing.T) {
	// 准备两个管子
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(time.Second*5)
		ch1 <- "test1"
	}()
	go func() {
		time.Sleep(time.Second*2)
		ch2 <- "test1"
	}()

	select {
	case s1 := <-ch1:
		t.Log("ch1 ", s1)
	case s2 := <-ch2:
		t.Log("ch2 ", s2)
	}
}

