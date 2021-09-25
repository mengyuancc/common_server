package tree

import "fmt"

// 中序遍历
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	fmt.Println(node.Value)
	node.Right.Traverse()
}