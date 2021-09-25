package main

import (
	"fmt"
	"imooc/tree"
)

type MyTreeNode struct {
	node *tree.Node
}

func (myNode *MyTreeNode) PostOrder() {
	fmt.Println(myNode.node.Left)
}

func main() {
	// var root TreeNode
	root := tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateTreeNode(2)
	// fmt.Print(root.Right)
	root.Traverse()

	a := MyTreeNode{&root}
	a.PostOrder()
}