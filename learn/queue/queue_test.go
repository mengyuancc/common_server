package queue

import (
	"fmt"
	"testing"
)

func TestQueue_Pop(t *testing.T) {
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	q.Pop()
	fmt.Println(q.IsEmpty())
	fmt.Println(q)
}

