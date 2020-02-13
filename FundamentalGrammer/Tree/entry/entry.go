package main

import (
	tree "GoLangIntro/FundamentalGrammer/Tree"
	"fmt"
)

func main() {
	fmt.Println("Go language Objective Oriented Programming with Tree as an example")
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	fmt.Printf("In-order traversal = ")
	root.Traverse()
	fmt.Println()
}
