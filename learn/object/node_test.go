package object

import (
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	user1 := &User{"mengyuan", 12}
	user1.SetName()
	fmt.Println(user1)

	root := Node{3, nil, nil}
	fmt.Println(root)
}