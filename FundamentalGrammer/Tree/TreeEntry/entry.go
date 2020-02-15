package main

import (
	tree "GoLangIntro/FundamentalGrammer/Tree"
	"fmt"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postTraverse() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postTraverse()
	right.postTraverse()
	myNode.node.Print()
}

func main() {
	fmt.Println("Go language Objective Oriented Programming with Tree as an example")
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	fmt.Printf("1. In-order traversal = ")
	root.Traverse()
	fmt.Println()

	fmt.Printf("2. Post-order traverse in entry main/package = ")
	myRoot := myTreeNode{&root}
	myRoot.postTraverse()
	fmt.Println("\n")

	fmt.Printf("3. Count number of nodes in the tree")
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Printf("Number of nodes = %d\n", nodeCount)
}
