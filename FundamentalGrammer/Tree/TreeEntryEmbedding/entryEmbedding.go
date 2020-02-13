package main

import (
	tree "GoLangIntro/FundamentalGrammer/Tree"
	"fmt"
)

type myTreeNode struct {
	*tree.Node
}

func (myNode *myTreeNode) postTraverse() {
	if myNode == nil || myNode.Node == nil {
		return
	}

	left := myTreeNode{myNode.Left}
	right := myTreeNode{myNode.Right}
	left.postTraverse()
	right.postTraverse()
	myNode.Print()
}

func main() {
	fmt.Println("Go language Objective Oriented Programming with Tree as an example")

	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	fmt.Printf("1. In-order traversal = ")
	root.Traverse()
	fmt.Println("\n")

	fmt.Printf("2. Post-order traverse in entry main/package = ")
	root.postTraverse()
	fmt.Println("\n")
}
