package main

import (
	"fmt"
	"imooc/tree"
)

type User struct {
	Name string
	Age int
}

func (u *User) formatName() {
	u.Name = "hi!" + u.Name
}

func main() {
	user1 := &User{"mengyuan", 12}
	user1.formatName()
	fmt.Println(user1)

	root := tree.Node{3, nil, nil}
	fmt.Println(root)
}
