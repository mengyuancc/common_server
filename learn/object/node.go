package object

import (
	"fmt"
)

type Node struct {
	Value int
	Left, Right *Node
}

// 创建节点
func CreateTreeNode(value int) *Node {
	return &Node{Value: value}
}

// 打印节点
func (node Node) print() {
	fmt.Print(node)
}

// 设置value 值传递
func (node Node) SetValue(value int) {
	node.Value = value
}
// 设置value 引用传递
func (node *Node) SetValuePro(value int) {
	node.Value = value
}

// 中序遍历
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	fmt.Println(node.Value)
	node.Right.Traverse()
}

type User struct {
	Name string
	Age int
}

// 设置实例的值
func (u *User) SetName() {
	u.Name = "update:" + u.Name
}
